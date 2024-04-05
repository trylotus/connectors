package main

import (
	"github.com/trylotus/connector"
	"github.com/trylotus/connectors/deribit"
	"github.com/trylotus/connectors/deribit/market"

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

	c := deribit.DeribitConnector{
		Connector: conn,
	}

	c.Start()
}
