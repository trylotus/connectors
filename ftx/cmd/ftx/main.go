package main

import (
	"github.com/nakji-network/connector"
	"github.com/nakji-network/connectors/ftx"
	"github.com/nakji-network/connectors/market"
)

func main() {

	connector := connector.NewConnector()

	// Register topic and protobuf type mappings
	connector.RegisterProtos(
		&market.OpenInterest{},
	)

	c := ftx.FTXConnector{
		Connector: connector,
	}

	c.Start()
}
