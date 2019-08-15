package bignum

import (
	"fmt"
	"github.com/jfixby/pin"
	"math/big"
	"testing"
)

func TestCalcSubsidy(t *testing.T) {
	subsidyBlocksNumber := int64(3)
	T := float64(7)
	testHeight := subsidyBlocksNumber
	totalSubsidy := &big.Rat{}
	totalSubsidy.SetFloat64(0)
	for blockNum := int64(0); blockNum <= testHeight; blockNum++ {
		sub := calcSubsidy(subsidyBlocksNumber, blockNum, T)
		//totalSubsidy += sub
		totalSubsidy = totalSubsidy.Add(totalSubsidy, sub)
		pin.D(fmt.Sprintf("%v", blockNum), sub)
	}
	pin.D(fmt.Sprintf("totalSubsidy"), totalSubsidy)
}

func calcSubsidy(subsidyBlocksNumber int64, height int64, totalSubsidy float64) *big.Rat {
	if height == 0 { //genesis block
		//return 0
	}
	TWO := big.NewInt(2)
	ZERO := big.NewRat(0, 1)
	H := big.NewInt(height)
	N := big.NewInt(subsidyBlocksNumber)

	lastBlockIndex := new(big.Int).SetInt64(subsidyBlocksNumber - 1)
	if H.Cmp(lastBlockIndex) > 0 {
		return ZERO
	}
	//endSubsidy := float64(0)               // 0 coins

	//return totalSubsidy * 2.0 * (lastBlockIndex - H) / (N * lastBlockIndex)
	H = H.Neg(H)                 // -H
	H = H.Add(lastBlockIndex, H) // (lastBlockIndex - H)
	H = H.Mul(TWO, H)            // 2.0 * (lastBlockIndex - H)
	N = N.Mul(N, lastBlockIndex) // (N * lastBlockIndex)

	subsidy := big.NewRat(1, 1)
	subsidy = subsidy.SetFrac(H, N) //  2.0 * (lastBlockIndex - H) / (N * lastBlockIndex)

	T := big.NewRat(1, 1)
	T = T.SetFloat64(totalSubsidy)
	T = T.Mul(T, subsidy) // totalSubsidy * 2.0 * (lastBlockIndex - H) / (N * lastBlockIndex)

	//float64Result, _ := T.Float64()
	return T
	//return totalSubsidy * 2.0 * (lastBlockIndex - H) / (N * lastBlockIndex)
}
