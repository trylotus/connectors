package main

import (
	"github.com/nakji-network/connector"
	"github.com/nakji-network/connectors/binance"
	"github.com/nakji-network/connectors/market"
)

func main() {

	connector := connector.NewConnector()

	// Register topic and protobuf type mappings
	connector.RegisterProtos(
		&market.OpenInterest{},
	)

	c := binance.BinanceConnector{
		Connector: connector,
	}

	c.Start()
}
