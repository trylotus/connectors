package main

import (
	"context"

	"github.com/nakji-network/connectors/bitcoin"
)

func main() {
	connector := bitcoin.NewConnector(func() {})

	// Register topic and protobuf type mappings
	connector.RegisterProtos(
		&bitcoin.Block{},
		&bitcoin.Transaction{},
	)

	connector.Start(context.Background())
}
