# 🐙 Connectors
This repo contains all the Lotus team supported connectors.

Github Actions is set up to build, run tests, and push docker image for each connector individually.

## 🏃‍♀️ To run
```
cd <connector_name>
go run cmd/<connector_name>/main.go
```

## 🚀 To deploy
For Lotus devs, ensure that the image was built successfully (check Github Actions) and pushed to our internal artifact registry, then update Helmfile values.yaml. If you are not a Lotus dev, please [reach out to our team](mailto:hello@trylotus.xyz).

## 🛠️ How to Build a Connector
There are **only TWO requirements** when it comes to writing a valid Lotus Source Connector:

1. Write [protobuf](https://developers.google.com/protocol-buffers) messages to [Lotus Message Queue (Kafka)](https://kafka.apache.org/) using [transactions](https://www.confluent.io/blog/transactions-apache-kafka/).
2. Manifest file that contains connector metadata (eg title, author, version)

There are no DSLs (domain specific languages) you have to learn, no weird structures you have to abide by, and no awkward terminology you have to remember.

***Just Send Data.***

Below, you will find an example of a connector. Alternatively, feel free to look around this repo to see how other connectors work.

### Example

This is an example of a barebones Ethereum connector which ingests Block and Transaction data into Lotus, written in Golang.

:::info
This example is currently 90% finished, with most key features present. Some of the libraries referenced in the Go files have not been published publicly yet. A fully functional source connector will be published in a public repo soon. 
:::

### Protobuf

First, a `protobuf` definition file needs to be created. Protobuf is the platform-neutral data format for almost all the *data in motion* within Lotus, whether it's into or out of the Lotus Message Queue. Make sure to use Version 3 of the Protobuf spec ([proto3 documentation](https://developers.google.com/protocol-buffers/docs/proto3)).

```protobuf title="/ethereum.proto"
syntax = "proto3";

import "google/protobuf/timestamp.proto";

package ethereum;

option go_package = "github.com/trylotus/connector/examples/ethereum_connector";

// to convert addresses from bytes to hex address, https://github.com/ethereum/go-ethereum/blob/4b2ff1457ac28fb2894485194e0e344e84c2bcd7/common/types.go#L210
message Transaction {
  google.protobuf.Timestamp Timestamp    = 1; //uint64
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
  google.protobuf.Timestamp Timestamp  = 1; //uint64
  string                    Hash       = 2;
  uint64                    Difficulty = 3; //bigint
  uint64                    Number     = 4; //bigint
  uint64                    GasLimit   = 5; // uint64
  uint64                    GasUsed    = 6; // uint64
  uint64                    Nonce      = 7; //[8]byte .Uint64()
}
```

The proto definitions can be compiled by the `protoc` compiler into any language it supports (currently 7) to enable seamless cross-platform/cross-service communication. For purposes of this example, follow the [Protobuf for Golang guide here](https://developers.google.com/protocol-buffers/docs/gotutorial) to get set up.

### Manifest
Create a file named `manifest.yaml` in the project root with the following fields.
```yaml title="/manifest.yaml"
package: ethereum
owner: allie
version: "0_0_0"
```

### Executable
The `cmd/ethereum/main.go` file prepares all the config variables and Kafka connections.

```go title="/cmd/ethereum/main.go"
// This connector ingests real time data from any EVM compatible chain.
package main

import (
	"fmt"

	"github.com/trylotus/connector/config"
	"github.com/trylotus/connector/examples/ethereum"
	. "github.com/trylotus/connector/kafkautils"
	"github.com/trylotus/connector/monitor"
	"github.com/trylotus/connector/manifest"
	"github.com/rs/zerolog/log"
)

var conf = config.GetConfig()

func init() {
	//conf.SetDefault("rpcs.ethereum.full", []string{"wss://mainnet.infura.io/ws/v3/"})
}

func main() {
	// Load manifest file and config variables
	m := manifest.Load()
	RPCURLs := conf.GetStringSlice("rpcs.ethereum.full")

	log.Info().
		Strs("RPCs", RPCURLs).
		Str("kafkaTxId", kafkaTransactionID).
		Msg("Starting EVM connector")

	// Initialize Kakfa Producer
	kp, err := NewProducer(conf.GetString("kafka.url"), m)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create new kafka producer")
	}

	connector := ethereum.EthereumConnector{
		KP:       kp,
		RPCURLs:  RPCURLs,
		Manifest: m,
	}

	connector.Start()
}
```

### Main connector code
The `ethereum.go` file connects to the data source (in this case, an Ethereum RPC). When data is received, it is cleaned and shaped into the Protobuf format defined above, and sent to the correct Topics in the Lotus Message Queue (Kafka).

```go title="/ethereum.go"
package ethereum

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/ethereum/go-ethereum/core/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/trylotus/connector/manifest"
	"github.com/rs/zerolog/log"

	. "github.com/trylotus/connector/kafkautils"
)

type EthereumConnector struct {
	KP       *Producer
	RPCURLs  []string
	Manifest manifest.Manifest
}

func (c EthereumConnector) Start() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// connect to Ethereum RPC websockets
	log.Info().Strs("url", c.RPCURLs).Msg("connecting to Ethereum RPC")
	client, err := ethclient.DialContext(context.Background(), c.RPCURLs)
	if err != nil {
		log.Fatal().Err(err).Msg("Ethereum RPC connection error")
	}

	// Subscribe to headers
  headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal().Err(err)
	}

	// kafkaKey examples: namespace.subject, eth.eth, arbitrum.arbitrum
	kafkaKey := []byte("eth.eth")

	defer client.Close()

	// Start kafka transactions
	c.KP.EnableTransactions()

	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)

			// Loop through this whenever a new header is received
			case header := <-headers:
				// Header doesn't contain all the block data so we have to call `BlockByHash` to retrieve it
				block, err := client.BlockByHash(context.Background(), header.Hash())
				if err != nil {
					log.Fatal(err)
				}

				PrintBlock(block)

				// Convert Geth.Block -> Block generated by protoc -> Protobuf -> kafka
				var blockData Block
				blockData.UnmarshalEthBlock(block) // defined elsewhere
				err = c.KP.WriteKafkaMessages(c.Manifest.Topic("block"), kafkaKey, &blockData)
				if err != nil {
					log.Error().Err(err).Msg("Kafka write proto")
				}

				// Convert Geth.Transaction -> Transaction generated by protoc -> Protobuf -> Kafka
				for _, tx := range block.Transactions() {
					txData := Transaction{}
					txData.UnmarshalEthTransaction(tx)
					txData.Timestamp = blockData.Timestamp // Timestamp isn't in the raw transaction from geth

					err := c.KP.WriteKafkaMessages(c.Manifest.Topic("tx"), kafkaKey, &txData)
					if err != nil {
						log.Error().Err(err).Msg("Kafka write proto")
					}
				}

				// Commit Kafka Transaction
				err = c.KP.CommitTransaction(nil)
				if err != nil {
					log.Error().Err(err).Msg("Processor: Failed to commit transaction")

					err = c.KP.AbortTransaction(nil)
					if err != nil {
						log.Fatal().Err(err).Msg("")
					}
				}
				// Start a new transaction
				err = c.KP.BeginTransaction()
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
			c.KP.Close()
			return
		}
	}
}

func PrintBlock(block *ethtypes.Block) {
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


## Advanced
In some cases, a source connector may use input data from the Lotus message queue instead of an external source. This reduces the need for every connector to make a new RPC request, at the cost of slightly increasing latency. For example:

```kroki imgType="mermaid" imgTitle="Aggregator connectors"
flowchart LR
  so[Ethereum] --> |eth source connector|MQ( Message Queue ) 
  MQ -->|aave source connector| MQ
  MQ -->|nft source connector| MQ
  MQ -->|your smart contract source connector| MQ

  style MQ fill:#9f9,stroke:#333,stroke-width:4px
```

This is a rare use case.

## Support
If you have any questions, please feel free to [reach out to our team](mailto:hello@trylotus.xyz) with any questions you may have.
