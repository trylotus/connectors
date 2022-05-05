package main

import (
	"context"

	"github.com/nakji-network/connectors/bitcoin"
)

func main() {
	connector := bitcoin.NewConnector(func() {})
	connector.Start(context.Background())
}
