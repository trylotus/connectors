package pancakeswapv2

import (
	"fmt"
	"math/big"
	"strings"
)

func tokenAmount(rawAmount *big.Int, decimals int64) *big.Rat {
	return new(big.Rat).SetFrac(rawAmount, new(big.Int).Exp(big.NewInt(10), big.NewInt(decimals), nil))
}

func floatString(n *big.Rat, prec int64) string {
	return removeTrailingZeros(n.FloatString(int(prec)))
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
