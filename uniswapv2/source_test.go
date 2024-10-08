package uniswapv2_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/trylotus/connectors/uniswapv2/erc20"
	"github.com/trylotus/go-connector/source/evm"
)

func Test1(t *testing.T) {

	client, err := evm.Dial("wss://eth-mainnet.g.alchemy.com/v2/MucMjmhWx62LBZgI1katEI0mjwA4c59z")
	if err != nil {
		panic(err)
	}

	tokenContract, err := erc20.NewErc20(common.HexToAddress("0x431ad2ff6a9C365805eBaD47Ee021148d6f7DBe0"), client)
	if err != nil {
		panic(err)
	}

	name, err := tokenContract.Name(nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(name)

	symbol, err := tokenContract.Symbol(nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(symbol)

	decimals, err := tokenContract.Decimals(nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(decimals)
}
