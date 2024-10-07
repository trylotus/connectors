package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/pflag"
	"github.com/trylotus/connectors/uniswapv2"
	"github.com/trylotus/connectors/uniswapv2/factory"
	"github.com/trylotus/connectors/uniswapv2/pair"
	"github.com/trylotus/go-connector"
	"github.com/trylotus/go-connector/common"
	"github.com/trylotus/go-connector/log"
	"github.com/trylotus/go-connector/source/evm"
)

const FactoryContractAddr = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"

func main() {
	_ = godotenv.Load()

	var (
		subscribe     bool
		subscribeFrom int64
		backfill      common.List
	)

	pflag.VarP(&backfill, "backfill", "b", "comma separated list of block numbers to backfill (optional)")
	pflag.BoolVarP(&subscribe, "subscribe", "s", true, "whether to subscribe to new blockchain events")
	pflag.Int64VarP(&subscribeFrom, "subscribe-from", "f", 0, "block number to subscribe from (default latest block)")

	pflag.Parse()

	ctx, cancel := common.ContextWithSignal(context.Background(), os.Interrupt)
	defer cancel()

	rpcUrl := os.Getenv("RPC_URL")

	log.Info().Str("url", rpcUrl).Msg("Connecting to RPC")

	client, err := evm.DialContext(ctx, rpcUrl)
	if err != nil {
		log.Fatal().Err(err).Str("url", rpcUrl).Msg("Failed to connect to RPC")
	}

	dataSource := os.Getenv("DATA_SOURCE")

	log.Info().Str("source", dataSource).Msg("Connecting to store")

	store, err := uniswapv2.NewStore(ctx, dataSource)
	if err != nil {
		log.Fatal().Err(err).Str("source", dataSource).Msg("Failed to create store")
	}

	source := uniswapv2.NewSource(client, store, FactoryContractAddr, uniswapv2.WithDefaultOptions())

	c := connector.NewConnector(source, connector.WithDefaultOptions())

	go c.RegisterDescriptor(ctx, factory.File_factory_factory_proto, pair.File_pair_pair_proto)

	if subscribe {
		c.Subscribe(ctx, subscribeFrom)
	}

	if !backfill.IsEmpty() {
		c.Backfill(ctx, backfill)
	}

	if err := c.Run(ctx); err != nil {
		log.Error().Err(err).Msg("Connector shutdown")
		return
	}

	log.Info().Msg("Connector shutdown")
}
