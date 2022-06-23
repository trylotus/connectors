package main

import (
	"github.com/nakji-network/connector"
	"github.com/nakji-network/connectors/bybit"
	"github.com/nakji-network/connectors/market"

	"github.com/rs/zerolog/log"
)

func main() {

	c, err := connector.NewProducerConnector()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to instantiate connector")
	}

	// Register topic and protobuf type mappings
	connector.RegisterProtos(
		&market.OpenInterest{},
	)

	c := bybit.BybitConnector{
		Connector: connector,
	}

	c.Start()
}
