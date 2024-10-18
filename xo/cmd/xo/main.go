package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
	"github.com/trylotus/connectors/xo/xo"
	"github.com/trylotus/go-connector"
	"github.com/trylotus/go-connector/common"
	"github.com/trylotus/go-connector/log"
	"github.com/trylotus/go-connector/manager"
	"github.com/trylotus/go-connector/registry"
	"github.com/trylotus/go-connector/source/evm"
)

const XOContractAddr = "0x84583e7d2d92d87d5b3bac850ab4bad37ae568e8"

func main() {
	_ = godotenv.Load()

	var (
		backfill  int64
		subscribe bool
	)

	pflag.Int64VarP(&backfill, "backfill", "b", 0, "block number to backfill to")
	pflag.BoolVarP(&subscribe, "subscribe", "s", true, "whether to subscribe to new blockchain events")

	pflag.Parse()

	contracts := []evm.SmartContract{
		xo.NewContract(XOContractAddr),
	}

	ctx, cancel := common.ContextWithSignal(context.Background(), os.Interrupt)
	defer cancel()

	rpcUrl := os.Getenv("RPC_URL")

	log.Info().Str("url", rpcUrl).Msg("Connecting to RPC")

	client, err := evm.DialContext(ctx, rpcUrl)
	if err != nil {
		log.Fatal().Err(err).Str("url", rpcUrl).Msg("Failed to connect to RPC")
	}

	c := connector.NewConnector(
		evm.NewSource(client, contracts),
		manager.NewManager(os.Getenv("MANAGER_URL")),
		registry.NewRegistry(os.Getenv("REGISTRY_URL")),
		connector.WithDefaultOptions(),
	)

	go c.RegisterDescriptor(ctx, xo.File_xo_xo_proto)

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
