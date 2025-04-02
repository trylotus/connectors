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

    
	"github.com/starbloom-ai/connectors/ether-1743613311/x/anon_bank_authority"
)

var ContractList = []evm.SmartContract{
    
        anon_bank_authority.NewContract(ethcommon.HexToAddress("0x36b0f8F5B789114C9Eb773822BFf69160C016700")),
}

func main() {
	_ = godotenv.Load()

	var (
		backfill  uint64
		subscribe bool
	)

	pflag.Uint64VarP(&backfill, "backfill", "b", 0, "block number to backfill to")
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
        
            anon_bank_authority.File_anon_bank_authority_contract_proto,
	)

	c.Run(ctx, backfill, subscribe)
}
