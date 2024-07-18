package main

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	connector "github.com/trylotus/connector/chain/ethereum"
	"github.com/trylotus/connector/common"
	"github.com/trylotus/connectors/cyber/xo"
)

const XOContractAddr = "0x84583e7d2d92d87d5b3bac850ab4bad37ae568e8"

func main() {
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

	contracts := []connector.SmartContract{
		xo.NewContract(XOContractAddr),
	}

	// Create a context that cancels upon receiving interrupt signal
	ctx, cancel := common.ContextWithSignal(context.Background(), os.Interrupt)
	defer cancel()

	c := connector.NewConnector(ctx, "cyber", contracts)

	c.Config.SetDefault("cyber.author", "lotus")
	c.Config.SetDefault("cyber.version", "0_0_0")
	c.Config.SetDefault("blockTime", 15*time.Second)
	c.Config.SetDefault("waitBlocks", 4)

	// Register topic and protobuf type mappings
	go c.RegisterDescriptor(ctx, xo.File_xo_xo_proto)

	c.Start(ctx, subscribe, fromBlock, toBlock, numBlocks)

	log.Info().Msg("connector shutdown")
}
