package uniswapv2

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

type Pair struct {
	Number  int64  `db:"number"`
	Address string `db:"address"`
	Token0  string `db:"token0"`
	Token1  string `db:"token1"`
}

type Token struct {
	Address  string `db:"address"`
	Name     string `db:"name"`
	Symbol   string `db:"symbol"`
	Decimals uint8  `db:"decimals"`
}

type Store struct {
	db         *sqlx.DB
	pairCache  *lru.ARCCache
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
		pairCache:  pairCache,
		tokenCache: tokenCache,
	}

	return &store, nil
}

func (s *Store) AllPairs(ctx context.Context) (<-chan *Pair, <-chan error) {
	pairCh := make(chan *Pair, 100)
	errCh := make(chan error, 1)

	go func() {
		defer close(pairCh)
		defer close(errCh)

		rows, err := s.db.QueryxContext(ctx, "SELECT * FROM pairs ORDER BY number")
		if err != nil {
			errCh <- err
			return
		}
		defer rows.Close()

		for rows.Next() {
			var pair Pair
			if err := rows.StructScan(&pair); err != nil {
				errCh <- err
				return
			}
			s.pairCache.Add(pair.Address, &pair)
			s.pairCache.Add(pair.Number, &pair)
			pairCh <- &pair
		}

		if err := rows.Err(); err != nil {
			errCh <- err
		}
	}()

	return pairCh, errCh
}

func (s *Store) AddPair(ctx context.Context, pair *Pair) error {
	s.pairCache.Add(pair.Address, pair)
	s.pairCache.Add(pair.Number, pair)

	_, err := s.db.NamedExecContext(ctx, "INSERT INTO pairs (number, address, token0, token1) VALUES (:number, :address, :token0, :token1) ON CONFLICT DO NOTHING", pair)

	return err
}

func (s *Store) GetPair(ctx context.Context, address string) (*Pair, error) {
	if pair, ok := s.pairCache.Get(address); ok {
		return pair.(*Pair), nil
	}

	var pair Pair
	err := s.db.GetContext(ctx, &pair, "SELECT * FROM pairs WHERE address = $1", address)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	s.pairCache.Add(pair.Address, &pair)
	s.pairCache.Add(pair.Number, &pair)

	return &pair, nil
}

func (s *Store) GetPairByNumber(ctx context.Context, number int64) (*Pair, error) {
	if pair, ok := s.pairCache.Get(number); ok {
		return pair.(*Pair), nil
	}

	var pair Pair
	err := s.db.GetContext(ctx, &pair, "SELECT * FROM pairs WHERE number = $1", number)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	s.pairCache.Add(pair.Address, &pair)
	s.pairCache.Add(pair.Number, &pair)

	return &pair, nil
}

func (s *Store) AddToken(ctx context.Context, token *Token) error {
	s.pairCache.Add(token.Address, token)

	_, err := s.db.NamedExecContext(ctx, "INSERT INTO tokens (address, name, symbol, decimals) VALUES (:address, :name, :symbol, :decimals) ON CONFLICT DO NOTHING", token)

	return err
}

func (s *Store) GetToken(ctx context.Context, address string) (*Token, error) {
	if token, ok := s.tokenCache.Get(address); ok {
		return token.(*Token), nil
	}

	var token Token
	err := s.db.GetContext(ctx, &token, "SELECT * FROM tokens WHERE address = $1", address)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	s.tokenCache.Add(token.Address, &token)

	return &token, nil
}
