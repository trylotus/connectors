package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	"github.com/trylotus/connectors/cyber/xo"
	"github.com/trylotus/go-connector"
	"github.com/trylotus/go-connector/common"
	"github.com/trylotus/go-connector/source/evm"
)

const XOContractAddr = "0x4D73AdB72bC3DD368966edD0f0b2148401A178E2"

func main() {
	_ = godotenv.Load()

	var (
		fromBlock uint64
		toBlock   uint64
		numBlocks uint64
		subscribe bool
	)

	pflag.Uint64VarP(&fromBlock, "from-block", "f", 0, "block number to start backfill from (optional)")
	pflag.Uint64VarP(&toBlock, "to-block", "t", 0, "block number to backfill to (optional)")
	pflag.Uint64VarP(&numBlocks, "num-blocks", "b", 0, "number of blocks to backfill (optional)")
	pflag.BoolVarP(&subscribe, "subscribe", "s", true, "whether to subscribe to new blockchain events")

	pflag.Parse()

	contracts := []evm.SmartContract{
		xo.NewContract(XOContractAddr),
	}

	source := evm.NewSource(os.Getenv("RPC_URL"), contracts)

	c := connector.NewConnector(source)

	// Create a context that cancels upon receiving interrupt signal
	ctx, cancel := common.ContextWithSignal(context.Background(), os.Interrupt)
	defer cancel()

	// Register topic and protobuf type mappings
	go c.RegisterDescriptor(ctx, xo.File_xo_xo_proto)

	if err := c.Connect(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to connect")
	}

	c.Start(ctx, subscribe, fromBlock, toBlock, numBlocks)

	log.Info().Msg("Connector shutdown")
}
