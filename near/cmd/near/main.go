package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	"github.com/trylotus/connector"
	"github.com/trylotus/connector/config"
	"github.com/trylotus/connector/kafkautils"
	"github.com/trylotus/connectors/near"
	"google.golang.org/protobuf/proto"
)

func main() {
	c, err := connector.NewConnector()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to instantiate connector")
	}

	c.Config.SetDefault("channelSize", 1000)
	c.Config.SetDefault("port", "3030")
	c.Config.SetDefault("backfillPort", "3031")
	c.Config.SetDefault("maxRetries", 5)

	pflag.Int64P("from-block", "f", 0, "block number to start backfill from (optional")
	pflag.Int64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")

	pflag.Parse()
	c.Config.BindPFlags(pflag.CommandLine)

	if err := validateFlags(c.Config); err != nil {
		log.Fatal().Err(err).Msg("input is not correct")
	}

	config := &near.Config{
		FromBlock:    c.Config.GetUint64("from-block"),
		NumBlocks:    c.Config.GetUint64("num-blocks"),
		Hostname:     "localhost",
		Port:         c.Config.GetString("port"),
		BackfillPort: c.Config.GetString("backfillPort"),
		MaxRetries:   c.Config.GetInt("maxRetries"),
		FctMsgTypes:  []proto.Message{&near.Block{}, &near.Transaction{}, &near.Receipt{}, &near.ExecutionOutcome{}},
		BfMsgTypes:   []proto.Message{&near.Block{}, &near.Transaction{}},
		ChannelSize:  c.Config.GetInt("channelSize"),
	}

	connector := near.New(c, config)

	// Register topic and protobuf type mappings
	connector.RegisterProtos(kafkautils.MsgTypeFct, connector.FctMsgTypes...)
	connector.RegisterProtos(kafkautils.MsgTypeBf, connector.BfMsgTypes...)

	connector.Start()
}

func validateFlags(conf config.IConfig) error {
	fromBlock := conf.GetInt64("from-block")
	numBlocks := conf.GetInt64("num-blocks")

	if fromBlock < 0 {
		return fmt.Errorf("backfill input value cannot be negative. from-block: %d", fromBlock)
	}

	if numBlocks < 0 {
		return fmt.Errorf("backfill input value cannot be negative. num-blocks: %d", numBlocks)
	}
	return nil
}
