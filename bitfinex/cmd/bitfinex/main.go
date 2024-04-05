package main

import (
	"github.com/trylotus/connector"
	"github.com/trylotus/connectors/bitfinex"
	"github.com/trylotus/connectors/bitfinex/market"

	"github.com/rs/zerolog/log"
)

func main() {

	conn, err := connector.NewConnector()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to instantiate connector")
	}

	// Register topic and protobuf type mappings
	conn.RegisterProtos(
		&market.OpenInterest{},
	)

	c := bitfinex.BitfinexConnector{
		Connector: conn,
	}

	c.Start()
}
