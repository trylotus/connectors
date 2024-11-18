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

    
	"github.com/trylotus/connectors/pancakeswapv2-1731910845/contracts/pancake_factory"
	"github.com/trylotus/connectors/pancakeswapv2-1731910845/contracts/tmc"
	"github.com/trylotus/connectors/pancakeswapv2-1731910845/contracts/bep20_usdt"
)

var ContractList = []evm.SmartContract{
    
        pancake_factory.NewContract(ethcommon.HexToAddress("0xca143ce32fe78f1f7019d7d551a6402fc5350c73")),
        tmc.NewContract(ethcommon.HexToAddress("0x558AbE16B071057FF9b99f7eED358B4F4708a384")),
        bep20_usdt.NewContract(ethcommon.HexToAddress("0x55d398326f99059fF775485246999027B3197955")),
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
        
            pancake_factory.File_pancake_factory_contract_proto,
            tmc.File_tmc_contract_proto,
            bep20_usdt.File_bep20_usdt_contract_proto,
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
