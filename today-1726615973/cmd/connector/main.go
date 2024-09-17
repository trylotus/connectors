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

    
	"github.com/trylotus/connectors/today-1726615973/contracts/ultra_light_node_v_2"
	"github.com/trylotus/connectors/today-1726615973/contracts/anon_bank_authority"
)

var ContractList = []evm.SmartContract{
    
        ultra_light_node_v_2.NewContract("0x4D73AdB72bC3DD368966edD0f0b2148401A178E2"),
        anon_bank_authority.NewContract("0x36b0f8F5B789114C9Eb773822BFf69160C016700"),
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

	source := evm.NewSource(ctx, os.Getenv("RPC_URL"), ContractList)

	c := connector.NewConnector(source, connector.WithDefaultOptions())

	go c.RegisterDescriptor(ctx,
        
            ultra_light_node_v_2.File_ultra_light_node_v_2_contract_proto,
            anon_bank_authority.File_anon_bank_authority_contract_proto,
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
