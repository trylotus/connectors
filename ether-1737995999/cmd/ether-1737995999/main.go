package main

import (
	"context"
	"os"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
	"github.com/trylotus/go-connector"
	"github.com/trylotus/go-connector/common"
	"github.com/trylotus/go-connector/log"
	"github.com/trylotus/go-connector/manager"
	"github.com/trylotus/go-connector/registry"
	"github.com/trylotus/go-connector/source/evm"

    
	"github.com/trylotus/connectors/ether-1737995999/contracts/ultra_light_node_v2"
	"github.com/trylotus/connectors/ether-1737995999/contracts/anon_bank_authority"
)

var ContractList = []evm.SmartContract{
    
        ultra_light_node_v2.NewContract(ethcommon.HexToAddress("0x4D73AdB72bC3DD368966edD0f0b2148401A178E2")),
        anon_bank_authority.NewContract(ethcommon.HexToAddress("0x36b0f8F5B789114C9Eb773822BFf69160C016700")),
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
		connector.WithDefaultOptions(),
	)

	go c.RegisterDescriptor(ctx,
        
            ultra_light_node_v2.File_ultra_light_node_v2_contract_proto,
            anon_bank_authority.File_anon_bank_authority_contract_proto,
	)

	if subscribe {
		c.Subscribe(ctx)
	}

	if backfill > 0 {
		c.Backfill(ctx, backfill)
	}

	if err := c.Run(ctx); err != nil {
		log.Error().Err(err).Msg("Connector shutdown")
		return
	}

	log.Info().Msg("Connector shutdown")
}
