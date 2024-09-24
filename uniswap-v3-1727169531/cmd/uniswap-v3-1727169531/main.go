package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/pflag"

	"github.com/trylotus/go-connector"
	"github.com/trylotus/go-connector/common"
	"github.com/trylotus/go-connector/log"
	"github.com/trylotus/go-connector/source/evm"

    
	"github.com/trylotus/connectors/uniswap-v3-1727169531/contracts/ultra_light_node_v2"
)

var ContractList = []evm.SmartContract{
    
        ultra_light_node_v2.NewContract("0x4D73AdB72bC3DD368966edD0f0b2148401A178E2"),
}

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

	source := evm.NewSource(client, ContractList)

	c := connector.NewConnector(source, connector.WithDefaultOptions())

	go c.RegisterDescriptor(ctx,
        
            ultra_light_node_v2.File_ultra_light_node_v2_contract_proto,
	)

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
