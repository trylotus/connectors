package main

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	connector "github.com/trylotus/connector/chain/ethereum"
	"github.com/trylotus/connector/common"
	"github.com/trylotus/connector/kafkautils"
	"github.com/trylotus/connectors/aave"
	"github.com/trylotus/connectors/aave/lendingpool"
	_ "go.uber.org/automaxprocs"
)

func main() {
	fromBlock := pflag.Uint64P("from-block", "f", 0, "block number to start backfill from (optional)")
	toBlock := pflag.Uint64P("to-block", "t", 0, "block number to backfill to (optional)")
	numBlocks := pflag.Uint64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")

	pflag.Parse()

	contracts := []connector.SmartContract{
		lendingpool.NewContract(aave.LendingPoolContractAddr),
	}

	// Create a context that cancels upon receiving interrupt signal
	ctx, cancel := common.ContextWithSignal(context.Background(), os.Interrupt)
	defer cancel()

	c := connector.NewConnector(ctx, "ethereum", contracts)

	c.Config.SetDefault("aave.author", "lotus")
	c.Config.SetDefault("aave.version", "0_0_0")
	c.Config.SetDefault("blockTime", 15*time.Second)
	c.Config.SetDefault("waitBlocks", 4)

	// Register topic and protobuf type mappings
	c.RegisterProtos(kafkautils.MsgTypeFct, aave.TopicTypes...)

	if *fromBlock > 0 || *toBlock > 0 || *numBlocks > 0 {
		go func() {
			msgs, err := c.BackfillWithOption(ctx, common.BackfillOption{
				FromBlock: *fromBlock,
				ToBlock:   *toBlock,
				NumBlocks: *numBlocks,
			})
			if err != nil {
				log.Error().Err(err).Msg("invalid backfill option")
				return
			}
			c.StreamProtoMessages(kafkautils.MsgTypeBf, msgs)
		}()
	}

	msgs := c.Subscribe(ctx)
	c.StreamProtoMessages(kafkautils.MsgTypeFct, msgs)

	log.Info().Msg("connector shutdown")
}
