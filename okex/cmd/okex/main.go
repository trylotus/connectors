package main

import (
	"github.com/trylotus/connector"
	"github.com/trylotus/connectors/okex"
	"github.com/trylotus/connectors/okex/market"

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

	c := okex.OkexConnector{
		Connector: conn,
	}

	c.Start()
}
