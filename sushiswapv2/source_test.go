package sushiswapv2

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
)

func TestXxx(t *testing.T) {
	// 0xE0B7927c4aF23765Cb51314A0E0521A9645F0E2A

	client, err := ethclient.Dial("wss://eth-mainnet.g.alchemy.com/v2/MucMjmhWx62LBZgI1katEI0mjwA4c59z")
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.BlockByNumber(context.Background(), big.NewInt(1000000000))
	t.Log(err)

	// tokenContract, err := erc20.NewErc20(common.HexToAddress("0xF967a43a63163c7a1Cedb46cc9b73dA0a2c20986"), client)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// decimals, err := tokenContract.Decimals(nil)
	// if err != nil {
	// 	t.Error(err)
	// } else {
	// 	t.Log(decimals)
	// }

	// name, err := tokenContract.Name(nil)
	// if err != nil {
	// 	t.Error(err)
	// } else {
	// 	t.Log(name)
	// }

	// symbol, err := tokenContract.Symbol(nil)
	// if err != nil {
	// 	t.Error(err)
	// } else {
	// 	t.Log(symbol)
	// }
}
