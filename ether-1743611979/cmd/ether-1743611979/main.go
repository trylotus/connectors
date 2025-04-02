package main

import (
	"context"
	"os"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/spf13/pflag"

	"github.com/starbloom-ai/go-connector"
	"github.com/starbloom-ai/go-connector/common"
	"github.com/starbloom-ai/go-connector/kafka"
	"github.com/starbloom-ai/go-connector/log"
	"github.com/starbloom-ai/go-connector/manager"
	"github.com/starbloom-ai/go-connector/registry"
	"github.com/starbloom-ai/go-connector/source/evm"

    
	"github.com/starbloom-ai/connectors/ether-1743611979/x/ultra_light_node_v2"
)

var ContractList = []evm.SmartContract{
    
        ultra_light_node_v2.NewContract(ethcommon.HexToAddress("0x4D73AdB72bC3DD368966edD0f0b2148401A178E2")),
}

func main() {
	_ = godotenv.Load()

	var (
		backfill  int64
		subscribe bool
	)

	pflag.Int64VarP(&backfill, "backfill", "b", 0, "block number to backfill to")
	pflag.BoolVarP(&subscribe, "subscribe", "s", true, "whether to subscribe to new blockchain events")

	pflag.Parse()

	ctx, cancel := common.ContextWithSignal(context.Background(), os.Interrupt)
	defer cancel()

	client, err := evm.Connect(ctx, os.Getenv("RPC_URL"), os.Getenv("CACHE_URL"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create RPC client")
	}

	c := connector.NewConnector(
		evm.NewSource(client, ContractList...),
		manager.NewManager(os.Getenv("MANAGER_URL")),
		registry.NewRegistry(os.Getenv("REGISTRY_URL")),
		kafka.NewProducerFromFile(common.EnvString("KAFKA_CONFIG", "kafka.json")),
		connector.WithDefaultOptions(),
	)

	go c.RegisterDescriptor(ctx,
        
            ultra_light_node_v2.File_ultra_light_node_v2_contract_proto,
	)

	c.Run(ctx, backfill, subscribe)
}
