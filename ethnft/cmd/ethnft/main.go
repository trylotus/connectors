// Code generated by connectorgen - Edit as necessary.

package main

import (
	"context"

	"github.com/nakji-network/connector"

	"github.com/nakji-network/connectors/ethnft"
)

func main() {
	c := connector.NewConnector()

	connector := ethnft.NewConnector(
		c,
	)

	connector.Start(context.Background()) //, c.Config.GetUint64("nft.backfillNumBlocks"))
}
