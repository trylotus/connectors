package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
	"github.com/trylotus/connectors/uniswap-v-2/contracts/defaultct"
	"github.com/trylotus/go-connector"
	"github.com/trylotus/go-connector/common"
	"github.com/trylotus/go-connector/log"
	"github.com/trylotus/go-connector/source/evm"
)

const ContractAddr = "0xCD6bcca48069f8588780dFA274960F15685aEe0e"

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

	contracts := []evm.SmartContract{
		defaultct.NewContract(ContractAddr),
	}

	ctx, cancel := common.ContextWithSignal(context.Background(), os.Interrupt)
	defer cancel()

	source := evm.NewSource(ctx, os.Getenv("RPC_URL"), contracts)

	c := connector.NewConnector(source, connector.WithDefaultOptions())

	go c.RegisterDescriptor(ctx, defaultct.File_defaultct_contract_proto)

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
