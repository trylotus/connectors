package main

import (
	"github.com/trylotus/connector"
	"github.com/trylotus/connector/kafkautils"
	"github.com/trylotus/connectors/bybit"
	"github.com/trylotus/connectors/bybit/market"

	"github.com/rs/zerolog/log"
)

func main() {

	conn, err := connector.NewConnector()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to instantiate connector")
	}

	// Register topic and protobuf type mappings
	conn.RegisterProtos(kafkautils.MsgTypeFct, &market.OpenInterest{})

	c := bybit.BybitConnector{
		Connector: conn,
	}

	c.Start()
}
