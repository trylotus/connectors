package main

import (
	"github.com/nakji-network/connector"
	"github.com/nakji-network/connectors/market"
	"github.com/nakji-network/connectors/okex"
)

func main() {

	connector := connector.NewConnector()

	// Register topic and protobuf type mappings
	connector.RegisterProtos(
		&market.OpenInterest{},
	)

	c := okex.OkexConnector{
		Connector: connector,
	}

	c.Start()
}
