package main

import (
	"context"
	"os"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/pflag"
	"github.com/trylotus/connectors/sushiswapv2"
	"github.com/trylotus/connectors/sushiswapv2/factory"
	"github.com/trylotus/connectors/sushiswapv2/pair"
	"github.com/trylotus/go-connector"
	"github.com/trylotus/go-connector/common"
	"github.com/trylotus/go-connector/log"
	"github.com/trylotus/go-connector/manager"
	"github.com/trylotus/go-connector/registry"
	"github.com/trylotus/go-connector/source/evm"
)

const FactoryContractAddr = "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac"

func main() {
	_ = godotenv.Load()

	var (
		backfill  int64
		subscribe bool
	)

	pflag.Int64VarP(&backfill, "backfill", "b", 0, "block number to backfill to")
	pflag.BoolVarP(&subscribe, "subscribe", "s", true, "whether to subscribe to new blockchain events")

	pflag.Parse()

	ctx, cancel := common.ContextWithSignal(context.Background(), os.Interrupt)
	defer cancel()

	client, err := evm.Connect(ctx, os.Getenv("RPC_URL"), os.Getenv("CACHE_URL"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create RPC client")
	}

	store, err := sushiswapv2.NewStore(ctx, os.Getenv("DATA_SOURCE"))
	if err != nil {
		log.Fatal().Err(err).Str("source", os.Getenv("DATA_SOURCE")).Msg("Failed to create store")
	}

	source := sushiswapv2.NewSource(
		client,
		store,
		ethcommon.HexToAddress(FactoryContractAddr),
		sushiswapv2.WithDefaultOptions(),
	)

	c := connector.NewConnector(
		source,
		manager.NewManager(os.Getenv("MANAGER_URL")),
		registry.NewRegistry(os.Getenv("REGISTRY_URL")),
		connector.WithDefaultOptions(),
	)

	go c.RegisterDescriptor(ctx, factory.File_factory_factory_proto, pair.File_pair_pair_proto)

	source.Init(ctx)

	if subscribe {
		c.Subscribe(ctx)
	}

	if backfill > 0 {
		c.Backfill(ctx, backfill)
	}

	if err := c.Run(ctx); err != nil {
		log.Error().Err(err).Msg("Connector shutdown")
		return
	}

	log.Info().Msg("Connector shutdown")
}
