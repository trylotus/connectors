package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
	"github.com/trylotus/connectors/monad"
	protoevm "github.com/trylotus/connectors/monad/evm"
	"github.com/trylotus/go-connector"
	"github.com/trylotus/go-connector/common"
	"github.com/trylotus/go-connector/log"
	"github.com/trylotus/go-connector/manager"
	"github.com/trylotus/go-connector/registry"
	"github.com/trylotus/go-connector/source/evm"
)

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

	client, err := evm.Connect(ctx, os.Getenv("RPC_URL"), "")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create RPC client")
	}

	c := connector.NewConnector(
		monad.NewSource(client),
		manager.NewManager(os.Getenv("MANAGER_URL")),
		registry.NewRegistry(os.Getenv("REGISTRY_URL")),
		connector.WithDefaultOptions(),
	)

	go c.RegisterDescriptor(ctx, protoevm.File_evm_evm_proto)

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
