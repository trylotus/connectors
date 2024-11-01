package sunswap

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestPoolList(t *testing.T) {
	var pools PoolList

	pools.Add(common.HexToAddress("0xD1D5A4c0eA98971894772Dcd6D2f1dc71083C44E"), 12369922)
	pools.Add(common.HexToAddress("0x1d42064Fc4Beb5F8aAF85F4617AE8b3b5B8Bd801"), 12369739)
	pools.Add(common.HexToAddress("0x6c6Bc977E13Df9b0de53b251522280BB72383700"), 12369760)
	pools.Add(common.HexToAddress("0x7BeA39867e4169DBe237d55C8242a8f2fcDcc387"), 12369811)
	pools.Add(common.HexToAddress("0xCBCdF9626bC03E24f779434178A73a0B4bad62eD"), 12369821)
	pools.Add(common.HexToAddress("0xC2e9F25Be6257c210d7Adf0D4Cd6E3E881ba25f8"), 12369854)
	pools.Add(common.HexToAddress("0x1d42064Fc4Beb5F8aAF85F4617AE8b3b5B8Bd801"), 12369739)
	pools.Add(common.HexToAddress("0x30EB5E15476E6a80F4F3cd8479749b4881DAB1b8"), 12370078)
	pools.Add(common.HexToAddress("0x7858E59e0C01EA06Df3aF3D20aC7B0003275D4Bf"), 12369863)
	pools.Add(common.HexToAddress("0x886072A44BDd944495eFF38AcE8cE75C1EacDAF6"), 12369879)
	pools.Add(common.HexToAddress("0xF83d5AaaB14507A53f97D3C18BDB52C4A62Efc40"), 12369901)
	pools.Add(common.HexToAddress("0xCBCdF9626bC03E24f779434178A73a0B4bad62eD"), 12369821)
	pools.Add(common.HexToAddress("0x6f48ECa74B38d2936B02ab603FF4e36A6C0E3A77"), 12370078)
	pools.Add(common.HexToAddress("0x30EB5E15476E6a80F4F3cd8479749b4881DAB1b8"), 12370078)

	pools.SortAndRemoveDuplicates()

	require.Equal(t, []int64{
		12369739,
		12369760,
		12369811,
		12369821,
		12369854,
		12369863,
		12369879,
		12369901,
		12369922,
		12370078,
		12370078,
	}, pools.BlockNumbers)

	require.Equal(t, []common.Address{
		common.HexToAddress("0x1d42064Fc4Beb5F8aAF85F4617AE8b3b5B8Bd801"),
		common.HexToAddress("0x6c6Bc977E13Df9b0de53b251522280BB72383700"),
		common.HexToAddress("0x7BeA39867e4169DBe237d55C8242a8f2fcDcc387"),
		common.HexToAddress("0xCBCdF9626bC03E24f779434178A73a0B4bad62eD"),
		common.HexToAddress("0xC2e9F25Be6257c210d7Adf0D4Cd6E3E881ba25f8"),
		common.HexToAddress("0x7858E59e0C01EA06Df3aF3D20aC7B0003275D4Bf"),
		common.HexToAddress("0x886072A44BDd944495eFF38AcE8cE75C1EacDAF6"),
		common.HexToAddress("0xF83d5AaaB14507A53f97D3C18BDB52C4A62Efc40"),
		common.HexToAddress("0xD1D5A4c0eA98971894772Dcd6D2f1dc71083C44E"),
		common.HexToAddress("0x30EB5E15476E6a80F4F3cd8479749b4881DAB1b8"),
		common.HexToAddress("0x6f48ECa74B38d2936B02ab603FF4e36A6C0E3A77"),
	}, pools.Addresses)

	require.Equal(t, []common.Address{
		common.HexToAddress("0x1d42064Fc4Beb5F8aAF85F4617AE8b3b5B8Bd801"),
		common.HexToAddress("0x6c6Bc977E13Df9b0de53b251522280BB72383700"),
		common.HexToAddress("0x7BeA39867e4169DBe237d55C8242a8f2fcDcc387"),
		common.HexToAddress("0xCBCdF9626bC03E24f779434178A73a0B4bad62eD"),
		common.HexToAddress("0xC2e9F25Be6257c210d7Adf0D4Cd6E3E881ba25f8"),
		common.HexToAddress("0x7858E59e0C01EA06Df3aF3D20aC7B0003275D4Bf"),
	}, pools.Search(12369863))
}
