package main

import (
	"github.com/nakji-network/connector"
	"github.com/nakji-network/connectors/bybit"
	"github.com/nakji-network/connectors/market"
)

func main() {

	connector := connector.NewConnector()

	// Register topic and protobuf type mappings
	connector.RegisterProtos(
		&market.OpenInterest{},
	)

	c := bybit.BybitConnector{
		Connector: connector,
	}

	c.Start()
}
