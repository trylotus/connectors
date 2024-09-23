package uniswapv2_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/trylotus/go-connector/source/evm"
	"github.com/trylotus/uniswapv2/pair"
)

func Test1(t *testing.T) {
	client, err := evm.Dial("wss://eth-mainnet.g.alchemy.com/v2/MucMjmhWx62LBZgI1katEI0mjwA4c59z")
	if err != nil {
		panic(err)
	}

	pair, err := pair.NewPair(common.HexToAddress("0x1c32cf7a84bfaba4040e5fbd6833f54f49d7ad9b"), client)
	if err != nil {
		panic(err)
	}

	name, err := pair.Name(nil)
	if err != nil {
		panic(err)
	}

	symbol, err := pair.Symbol(nil)
	if err != nil {
		panic(err)
	}

	decimals, err := pair.Decimals(nil)
	if err != nil {
		panic(err)
	}

	t.Log(name, symbol, decimals)
}
