package sunswap

import (
	"sort"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

type PoolList struct {
	Addresses    []ethcommon.Address
	BlockNumbers []int64
}

func (pl *PoolList) Add(address ethcommon.Address, blockNumber int64) {
	pl.Addresses = append(pl.Addresses, address)
	pl.BlockNumbers = append(pl.BlockNumbers, blockNumber)
}

func (pl *PoolList) SortAndRemoveDuplicates() {
	if len(pl.BlockNumbers) == 0 || len(pl.Addresses) == 0 {
		return
	}

	if len(pl.BlockNumbers) <= 1 {
		return
	}

	indices := make([]int, len(pl.BlockNumbers))
	for i := range indices {
		indices[i] = i
	}

	sort.Slice(indices, func(i, j int) bool {
		b1 := pl.BlockNumbers[indices[i]]
		b2 := pl.BlockNumbers[indices[j]]
		a1 := pl.Addresses[indices[i]]
		a2 := pl.Addresses[indices[j]]
		return (b1 < b2) || ((b1 == b2) && (a1.Cmp(a2) < 0))
	})

	sortedAddresses := make([]ethcommon.Address, len(pl.Addresses))
	sortedBlockNumbers := make([]int64, len(pl.BlockNumbers))

	for i, index := range indices {
		sortedAddresses[i] = pl.Addresses[index]
		sortedBlockNumbers[i] = pl.BlockNumbers[index]
	}

	pl.Addresses = sortedAddresses
	pl.BlockNumbers = sortedBlockNumbers

	i := 0
	for j := 1; j < len(pl.Addresses); j++ {
		if pl.Addresses[i] != pl.Addresses[j] {
			i++
			pl.Addresses[i] = pl.Addresses[j]
			pl.BlockNumbers[i] = pl.BlockNumbers[j]
		}
	}
	pl.Addresses = pl.Addresses[:i+1]
	pl.BlockNumbers = pl.BlockNumbers[:i+1]
}

// Search returns all pools that are created before blockNumber + 1
func (pl *PoolList) Search(blockNumber int64) []ethcommon.Address {
	index := sort.Search(len(pl.BlockNumbers), func(i int) bool {
		return pl.BlockNumbers[i] > blockNumber
	})

	return pl.Addresses[:index]
}

func (pl *PoolList) Len() int {
	return len(pl.Addresses)
}
