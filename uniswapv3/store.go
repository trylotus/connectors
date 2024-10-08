package uniswapv3

import (
	"context"
	"database/sql"
	"errors"

	lru "github.com/hashicorp/golang-lru"
	"github.com/jmoiron/sqlx"
)

const (
	pairCacheSize  = 10000
	tokenCacheSize = 20000
)

type Pool struct {
	Address     string `db:"address"`
	Token0      string `db:"token0"`
	Token1      string `db:"token1"`
	Fee         int64  `db:"fee"`
	TickSpacing int64  `db:"tick_spacing"`
}

type Token struct {
	Address  string `db:"address"`
	Name     string `db:"name"`
	Symbol   string `db:"symbol"`
	Decimals uint8  `db:"decimals"`
}

type Store struct {
	db         *sqlx.DB
	poolCache  *lru.ARCCache
	tokenCache *lru.ARCCache
}

func NewStore(ctx context.Context, dataSource string) (*Store, error) {
	db, err := sqlx.ConnectContext(ctx, "postgres", dataSource)
	if err != nil {
		return nil, err
	}
	pairCache, err := lru.NewARC(pairCacheSize)
	if err != nil {
		return nil, err
	}
	tokenCache, err := lru.NewARC(tokenCacheSize)
	if err != nil {
		return nil, err
	}

	store := Store{
		db:         db,
		poolCache:  pairCache,
		tokenCache: tokenCache,
	}

	return &store, nil
}

func (s *Store) GetScannedBlock(ctx context.Context) (int64, error) {
	var number int64
	if err := s.db.GetContext(ctx, &number, "SELECT number from v3_pools_scanned_block where id = 1"); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}

	return number, nil
}

func (s *Store) SetScannedBlock(ctx context.Context, number int64) error {
	_, err := s.db.ExecContext(ctx, "INSERT INTO v3_pools_scanned_block (id, number) VALUES (1, $1) ON CONFLICT (id) DO UPDATE SET number = EXCLUDED.number", number)

	return err
}

func (s *Store) AllPools(ctx context.Context) (<-chan *Pool, <-chan error) {
	poolCh := make(chan *Pool, 100)
	errCh := make(chan error, 1)

	go func() {
		defer close(poolCh)
		defer close(errCh)

		rows, err := s.db.QueryxContext(ctx, "SELECT * FROM v3_pools")
		if err != nil {
			errCh <- err
			return
		}
		defer rows.Close()

		for rows.Next() {
			var pool Pool
			if err := rows.StructScan(&pool); err != nil {
				errCh <- err
				return
			}
			s.poolCache.Add(pool.Address, &pool)
			poolCh <- &pool
		}

		if err := rows.Err(); err != nil {
			errCh <- err
		}
	}()

	return poolCh, errCh
}

func (s *Store) AddPool(ctx context.Context, pool *Pool) error {
	s.poolCache.Add(pool.Address, pool)

	_, err := s.db.NamedExecContext(ctx, "INSERT INTO v3_pools (address, token0, token1, fee, tick_spacing) VALUES (:address, :token0, :token1, :fee, :tick_spacing) ON CONFLICT DO NOTHING", pool)

	return err
}

func (s *Store) GetPool(ctx context.Context, address string) (*Pool, error) {
	if pair, ok := s.poolCache.Get(address); ok {
		return pair.(*Pool), nil
	}

	var pool Pool
	err := s.db.GetContext(ctx, &pool, "SELECT * FROM v3_pools WHERE address = $1", address)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	s.poolCache.Add(pool.Address, &pool)

	return &pool, nil
}

func (s *Store) AddToken(ctx context.Context, token *Token) error {
	s.tokenCache.Add(token.Address, token)

	_, err := s.db.NamedExecContext(ctx, "INSERT INTO erc20_tokens (address, name, symbol, decimals) VALUES (:address, :name, :symbol, :decimals) ON CONFLICT DO NOTHING", token)

	return err
}

func (s *Store) GetToken(ctx context.Context, address string) (*Token, error) {
	if token, ok := s.tokenCache.Get(address); ok {
		return token.(*Token), nil
	}

	var token Token
	err := s.db.GetContext(ctx, &token, "SELECT * FROM erc20_tokens WHERE address = $1", address)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	s.tokenCache.Add(token.Address, &token)

	return &token, nil
}
