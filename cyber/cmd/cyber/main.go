package main

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	connector "github.com/trylotus/connector/chain/ethereum"
	"github.com/trylotus/connector/common"
	"github.com/trylotus/connector/kafkautils"
	"github.com/trylotus/connectors/cyber"
	"github.com/trylotus/connectors/cyber/xo"
)

func main() {
	fromBlock := pflag.Uint64P("from-block", "f", 0, "block number to start backfill from (optional)")
	toBlock := pflag.Uint64P("to-block", "t", 0, "block number to backfill to (optional)")
	numBlocks := pflag.Uint64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")
	subscribe := pflag.BoolP("subscribe", "s", true, "whether to subscribe to new blockchain events")

	pflag.Parse()

	contracts := []connector.SmartContract{
		xo.NewContract(cyber.XOContractAddr),
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
	c.RegisterProtos(kafkautils.MsgTypeFct, cyber.TopicTypes...)

	var wg sync.WaitGroup

	if *fromBlock > 0 || *toBlock > 0 || *numBlocks > 0 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			opt := common.BackfillOption{
				FromBlock: *fromBlock,
				ToBlock:   *toBlock,
				NumBlocks: *numBlocks,
			}
			msgs := c.BackfillWithOption(ctx, opt)
			c.StreamProtoMessages(ctx, kafkautils.MsgTypeBf, msgs)
		}()
	}

	if *subscribe {
		wg.Add(1)

		go func() {
			defer wg.Done()

			msgs := c.Subscribe(ctx)
			c.StreamProtoMessages(ctx, kafkautils.MsgTypeFct, msgs)
		}()
	}

	wg.Wait()

	log.Info().Msg("connector shutdown")
}
