package uniswapv3

import (
	"fmt"
	"math/big"
	"strings"
)

func tokenAmount(rawAmount *big.Int, decimals uint8) *big.Rat {
	return new(big.Rat).SetFrac(rawAmount, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
}

func floatString(n *big.Rat) string {
	return removeTrailingZeros(n.FloatString(18))
}

func removeTrailingZeros(number string) string {
	decimalIndex := strings.Index(number, ".")
	if decimalIndex == -1 {
		return number
	}

	integralPart := number[:decimalIndex]
	factionalPart := strings.TrimRight(number[decimalIndex+1:], "0")

	if factionalPart == "" {
		return integralPart
	} else {
		return fmt.Sprintf("%s.%s", integralPart, factionalPart)
	}
}
