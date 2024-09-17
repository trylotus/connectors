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

    
	"github.com/trylotus/connectors/hoangtiencong-1726598027/contracts/contract_0"
    
	"github.com/trylotus/connectors/hoangtiencong-1726598027/contracts/contract_1"
    
)

var ContractList = []evm.SmartContract{
    
        contract_0.NewContract("0xDC3D8318Fbaec2de49281843f5bba22e78338146"),
    
        contract_1.NewContract("0x2C512A4C0256A2b2146f4658075DBb9Dc8744a80"),
    
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
        
            contract_0.File_contract_0_contract_proto,
        
            contract_1.File_contract_1_contract_proto,
        
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
