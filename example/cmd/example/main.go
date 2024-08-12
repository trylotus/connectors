package main

import (
	"context"
	"crypto/rand"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/trylotus/connector"
	"github.com/trylotus/connector/common"
	"github.com/trylotus/connector/kafkautils"
	"github.com/trylotus/connectors/example/example"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const blockTime = 1 * time.Second
const txPerBlock = 10
const eventPerTx = 10

func main() {
	// Create a context that cancels upon receiving interrupt signal
	ctx, cancel := common.ContextWithSignal(context.Background(), os.Interrupt)
	defer cancel()

	c, err := connector.NewConnector()
	if err != nil {
		panic(err)
	}

	c.Config.SetDefault("example.author", "lotus")
	c.Config.SetDefault("example.version", "0_0_0")
	c.Config.SetDefault("blockTime", 15*time.Second)
	c.Config.SetDefault("waitBlocks", 4)

	// Register topic and protobuf type mappings
	go c.RegisterDescriptor(ctx, example.File_example_message_proto)

	ch := make(chan proto.Message, 1000)

	go func() {
		blockNumber := uint64(1)

		t := time.NewTicker(blockTime)

		for {
			now := time.Now()

			for i := 0; i < txPerBlock; i++ {
				txHash := make([]byte, 32)
				rand.Read(txHash)

				data := make(([]byte), 20)
				rand.Read(data)

				for index := uint64(1); index <= eventPerTx; index++ {
					msg := example.Message{
						Ts:          timestamppb.New(now),
						BlockNumber: blockNumber,
						TxHash:      txHash,
						Index:       index,
						Data:        data,
					}

					ch <- &msg
				}
			}

			<-t.C

			blockNumber++
		}
	}()

	c.StreamProtoMessages(ctx, kafkautils.MsgTypeFct, ch)

	log.Info().Msg("connector shutdown")
}
