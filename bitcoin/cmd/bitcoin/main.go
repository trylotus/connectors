package main

import (
	"context"

	"github.com/trylotus/connectors/bitcoin"

	_ "go.uber.org/automaxprocs"
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
