package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	connector "github.com/trylotus/connector/chain/ethereum"
	"github.com/trylotus/connector/kafkautils"
	"github.com/trylotus/connectors/aave"
	"github.com/trylotus/connectors/aave/lendingpool"
	_ "go.uber.org/automaxprocs"
)

func main() {
	fromBlock := pflag.Uint64P("from-block", "f", 0, "block number to start backfill from (optional)")
	numBlocks := pflag.Uint64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")

	pflag.Parse()

	contracts := []connector.SmartContract{
		lendingpool.NewContract(aave.LendingPoolContractAddr),
	}

	// Create a context that cancels on interrupt signal
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)

		<-interrupt
		log.Info().Msg("connector shutdown")
		cancel()
	}()

	c := connector.NewConnector(ctx, "ethereum", contracts)

	c.Config.SetDefault("aave.author", "lotus")
	c.Config.SetDefault("aave.version", "0_0_0")
	c.Config.SetDefault("blockTime", 15*time.Second)
	c.Config.SetDefault("waitBlocks", 4)

	// Register topic and protobuf type mappings
	c.RegisterProtos(kafkautils.MsgTypeFct, aave.TopicTypes...)

	if *fromBlock > 0 {
		go func() {
			msgs := c.Backfill(ctx, *fromBlock, *fromBlock+*numBlocks)
			c.StreamProtoMessages(kafkautils.MsgTypeBf, msgs)
		}()
	}

	msgs := c.Subscribe(ctx)
	c.StreamProtoMessages(kafkautils.MsgTypeFct, msgs)
}
