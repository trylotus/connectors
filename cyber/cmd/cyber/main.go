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
	"github.com/trylotus/connectors/cyber"
	"github.com/trylotus/connectors/cyber/TransparentUpgradeableProxy"
)

func main() {
	fromBlock := pflag.Uint64P("from-block", "f", 0, "block number to start backfill from (optional)")
	toBlock := pflag.Uint64P("to-block", "t", 0, "block number to backfill to (optional)")
	numBlocks := pflag.Uint64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")
	subscribe := pflag.BoolP("subscribe", "s", true, "whether to subscribe to new blockchain events (default: true)")

	pflag.Parse()

	contracts := []connector.SmartContract{
		TransparentUpgradeableProxy.NewContract(cyber.TransparentUpgradeableProxyContractAddr),
	}

	// Create a context that cancels upon receiving interrupt signal
	ctx, cancel := common.ContextWithSignal(context.Background(), os.Interrupt)
	defer cancel()

	c := connector.NewConnector(ctx, "cyber", contracts)

	c.Config.SetDefault("cyber.author", "cyber")
	c.Config.SetDefault("cyber.version", "0_0_0")
	c.Config.SetDefault("blockTime", 15*time.Second)
	c.Config.SetDefault("waitBlocks", 4)

	// Register topic and protobuf type mappings
	c.RegisterProtos(kafkautils.MsgTypeFct, cyber.TopicTypes...)

	if *fromBlock > 0 || *toBlock > 0 || *numBlocks > 0 {
		go func() {
			opt := common.BackfillOption{
				FromBlock: *fromBlock,
				ToBlock:   *toBlock,
				NumBlocks: *numBlocks,
			}
			msgs, err := c.BackfillWithOption(ctx, opt)
			if err != nil {
				log.Error().Err(err).Msg("invalid backfill option")
				return
			}
			if err := c.StreamProtoMessages(ctx, kafkautils.MsgTypeBf, msgs); err != nil {
				log.Error().Err(err).Msg("failed to send historical events to kafka")
			}
		}()
	}

	if *subscribe {
		msgs := c.Subscribe(ctx)
		if err := c.StreamProtoMessages(ctx, kafkautils.MsgTypeFct, msgs); err != nil {
			log.Error().Err(err).Msg("failed to send live events to kafka")
		}
	}

	log.Info().Msg("connector shutdown")
}
