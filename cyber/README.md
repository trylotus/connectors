# 🤔💭 What is Lotus?
Navigating on-chain data is complex and slow, making it a resource-heavy task for many projects. Current market solutions don't fully meet the need for specialized data infrastructure, leading many dApps to build their own due to three key issues: fragmented data sources, lack of "real-time" data, and unreliable uptime.

Lotus solves the above problems by featuring a modular design that allows for incremental updates and customization to meet specific data requirements. At its core, Lotus **connectors** are modules that facilitate data movement and interaction within the pipeline, linking data sources, message queues, and databases. This modularity makes the Lotus system adaptable, upgradable, and scalable, catering to the diverse data needs of decentralized applications (dApps).

ajisdfjdsaf asifjasd ioajsdfioadsjf
adsfjiadsfdsoafjsf
husadhasdhsahiasd

dsafijadsf
asdfiadsf
jisadfjdsf
asdfji

final check
final check


## 🔗 Connectors
Lotus connectors are essential for managing data flow in Lotus's data pipeline. They serve as the link between data sources, message queues, and databases, enhancing Lotus's modularity, upgradeability, and scalability.

Each connector performs three main functions: data **input**, data **transformation**, and data **output**, functioning similarly to a stream processor or a [streaming ETL](https://hazelcast.com/glossary/streaming-etl/).

By introducing a new connector to Lotus, you can achieve various tasks, such as:

- Data ingestion
- Data transformation and augmentation
- Data indexing and streaming

These connectors enable all Lotus data consumers to access and query data efficiently.

This repo contains all the Lotus team supported connectors.

Github Actions is set up to build, run tests, and push docker image for each connector individually.

## 🏃‍♀️ To run
```
cd <connector_name>
go run cmd/<connector_name>/main.go
```

## 🚀 To deploy
For Lotus devs, ensure that the image was built successfully (check Github Actions) and pushed to our internal artifact registry, then update Helmfile values.yaml. 

If you would like to build your own connector, please [reach out to our team](mailto:hello@trylotus.xyz).

## 🛠️ Source Connector Example

To create a valid Lotus Source Connector, you need to fulfill just two essential requirements:

1. Utilize [protobuf (proto3)](https://developers.google.com/protocol-buffers) to send messages to [Lotus Message Queue (Kafka)](https://kafka.apache.org/) within [transactions](https://www.confluent.io/blog/transactions-apache-kafka/).
2. Include a manifest file detailing the connector's metadata (e.g., title, author, version).

There are no need to learn DSLs (domain specific languages), adhere to complex structures, or remember unfamiliar terms. **Simply focus on data transmission.**

### Example

Explore a minimal Ethereum connector example for Lotus, developed in Golang. This connector processes Block and Transaction data.

Feel free to dive straight into different connectors in this repo if you prefer to learn by doing :)

:::info
This is beta software! Bug reports and suggestions are highly encouraged :)
:::

### Prerequisites

- [Golang 1.20+](https://go.dev/doc/install)
- [Kafka](https://github.com/wurstmeister/kafka-docker)
- [Protobuf (proto3)](https://developers.google.com/protocol-buffers/docs/gotutorial)

### Protobuf

Start by creating a protobuf definition file. Protobuf, particularly Version 3 ([proto3 documentation](https://developers.google.com/protocol-buffers/docs/proto3)) is crucial for Lotus's data handling. Protobuf is the platform-neutral data format for almost all the *data in motion* within Lotus, whether it's into or out of the Lotus Message Queue.

```protobuf title="/ethereum.proto"
syntax = "proto3";

import "google/protobuf/timestamp.proto";

package ethereum;

option go_package = "github.com/trylotus/connector/examples/ethereum_connector";

// to convert addresses from bytes to hex address, https://github.com/ethereum/go-ethereum/blob/4b2ff1457ac28fb2894485194e0e344e84c2bcd7/common/types.go#L210
message Transaction {
  google.protobuf.Timestamp ts    = 1; //uint64
  bytes                     From         = 2;
  string                    Hash         = 3;
  double                    Size         = 4;
  uint64                    AccountNonce = 5; // uint64
  uint64                    Price        = 6; // big.int
  uint64                    GasLimit     = 7; // uint64
  bytes                     Recipient    = 8;
  uint64                    Amount       = 9; // big.int
  bytes                     Payload      = 10;
  uint64                    V            = 11; // big.int
  uint64                    R            = 12; // big.int
  uint64                    S            = 13; // big.int
}

message Block {
  google.protobuf.Timestamp ts  = 1; //uint64
  string                    Hash       = 2;
  uint64                    Difficulty = 3; //bigint
  uint64                    Number     = 4; //bigint
  uint64                    GasLimit   = 5; // uint64
  uint64                    GasUsed    = 6; // uint64
  uint64                    Nonce      = 7; //[8]byte .Uint64()
}
```

Use the `protoc` compiler to convert these definitions into any language it supports (currently 7) to enable seamless cross-platform/cross-service communication. For purposes of this example, we will use the [Protobuf for Golang](https://developers.google.com/protocol-buffers/docs/gotutorial).

### Manifest
Create a `manifest.yaml`` in the project root with basic metadata.

```yaml title="/manifest.yaml"
package: ethereum
owner: Lotus
version: 0.0.0
```

### Executable
Configure Kafka connections and other settings in `'cmd/ethereum/main.go'`.

```go title="/cmd/ethereum/main.go"
// This connector ingests real time data from any EVM compatible chain.
package main

import (
	"strings"

	"github.com/trylotus/connector"
	"github.com/trylotus/connector/examples/ethereum"
)

func init() {
	conf.SetDefault("rpcs.ethereum.full", []string{"wss://mainnet.infura.io/ws/v3/<api_key>"})
}

func main() {
	c := connector.NewConnector()

	// Get config variables using functions from Viper (https://pkg.go.dev/github.com/spf13/viper#readme-getting-values-from-viper)
	RPCURLs := c.Config.GetStringSlice("rpcs.ethereum.full")

	// For the purposes of this example, we'll just grab one of the websocket RPCs
	var RPCURL string
	for _, u := range RPCURLs {
		if strings.HasPrefix(u, "ws") {
			RPCURL = u
			break
		}
	}

	ethConnector := ethereum.EthereumConnector{
		Connector: c,
		RPCURL:   RPCURL,
	}

	ethConnector.Start()
}
```

### Main connector code
`ethereum.go` manages data from the Ethereum RPC, formats it using Protobuf format as we defined it above, and sends it to Lotus Message Queue (Kafka).

```go title="/ethereum.go"
// ethereum package follows https://goethereumbook.org/block-subscribe/ to
// subscribe to new Blocks and Transactions and writes the results to Lotus.
package ethereum

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/trylotus/connector"
	"github.com/rs/zerolog/log"
)

type EthereumConnector struct {
	*connector.Connector // embed Lotus connector.Connector into your custom connector to get access to all its methods

	RPCURL string
}

func (c *EthereumConnector) Start() {
	// Listen for interrupt in order to cleanly close connections later
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// connect to Ethereum RPC websockets
	log.Info().Str("url", c.RPCURL).Msg("connecting to Ethereum RPC")
	client, err := ethclient.DialContext(context.Background(), c.RPCURL)
	if err != nil {
		log.Fatal().Err(err).Msg("Ethereum RPC connection error")
	}
	defer client.Close()

	// Subscribe to headers
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal().Err(err)
	}

	// Main loop to process errors and headers
	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatal().Err(err)
			case header := <-headers:
				block, err := client.BlockByHash(context.Background(), header.Hash())
				if err != nil {
					log.Fatal().Err(err).Msg("BlockByHash error")
				}

				PrintBlock(block)

				// EthBlock -> Block -> Protobuf -> kafka
				var blockData Block
				blockData.UnmarshalEthBlock(block)
				err = c.ProduceMessage("ethereum", "ethereum", &blockData)
				if err != nil {
					log.Error().Err(err).Msg("Kafka write proto")
				}

				// EthTransaction -> Transaction -> Protobuf -> Kafka
				for _, tx := range block.Transactions() {
					txData := Transaction{}
					txData.UnmarshalEthTransaction(tx)
					txData.Ts = blockData.Ts // Timestamp isn't in the raw transaction from geth

					err = c.ProduceMessage("ethereum", "ethereum", &txData)
					if err != nil {
						log.Error().Err(err).Msg("Kafka write proto")
					}
				}

				// Commit Kafka Transaction
				err = c.Producer.CommitTransaction(nil)
				if err != nil {
					log.Error().Err(err).Msg("Processor: Failed to commit transaction")

					err = c.Producer.AbortTransaction(nil)
					if err != nil {
						log.Fatal().Err(err).Msg("")
					}
				}
				// Start a new transaction
				err = c.BeginTransaction()
				if err != nil {
					log.Fatal().Err(err).Msg("")
				}

			}
		}
	}()

	for {
		select {
		case <-interrupt:
			log.Debug().Msg("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			client.Close()
			c.Producer.Close()
			return
		}
	}
}

func PrintBlock(block *types.Block) {
	fmt.Printf("hash: %s\n", block.Hash().Hex())       // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
	fmt.Printf("num: %v\n", block.Number().Uint64())   // 3477413
	fmt.Printf("time: %v\n", block.Time())             // 1529525947
	fmt.Printf("nonce: %v\n", block.Nonce())           // 130524141876765836
	fmt.Printf("#tx: %v\n", len(block.Transactions())) // 7
	fmt.Printf("gaslim: %v\n", block.GasLimit())       // 1529525947
	fmt.Printf("gasuse: %v\n", block.GasUsed())        // 1529525947
	fmt.Printf("diff: %v\n", block.Difficulty())       // 1529525947
}
```

This is the basic formation of a Lotus connector. For full examples, check out the other connectors available on this repo.

## Support

If you have any questions or comments, please [reach out to our team](mailto:hello@trylotus.xyz).
