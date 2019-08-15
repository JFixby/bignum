package bignum

import (
	"fmt"
	"testing"
)

func TestCalcSubsidy(t *testing.T) {
	subsidyBlocksNumber := int64(16)
	targetTotalSubsidy := float64(1)
	Engine = Float64Engine{}
	resultFloat64 := testCalcSubsidy(subsidyBlocksNumber, targetTotalSubsidy, 1).ToFloat64()
	Engine = BigIntEngine{}
	resultBigFloat := testCalcSubsidy(subsidyBlocksNumber, targetTotalSubsidy, 1).ToFloat64()

	if resultFloat64 != (resultBigFloat) {
		t.Fatalf("mismatched total subsidy -- \n got %v, \nwant %v", resultFloat64, resultBigFloat)
	}
}

func testCalcSubsidy(subsidyBlocksNumber int64, targetTotalSubsidy float64, printIterations int64) BigNum {
	testHeight := subsidyBlocksNumber
	totalSubsidy := NewBigNum(0)
	for blockNum := int64(0); blockNum <= testHeight; blockNum++ {
		sub := calcSubsidy(subsidyBlocksNumber, blockNum, targetTotalSubsidy)
		//totalSubsidy += sub
		totalSubsidy = totalSubsidy.Add(totalSubsidy, sub)
		if blockNum%printIterations == 0 {
			blockNumPad := fmt.Sprintf("%2v", blockNum)
			subPad := fmt.Sprintf("%-20v", sub.ToFloat64())
			totalSubsidyPad := fmt.Sprintf("%-20v", totalSubsidy.ToFloat64())
			fmt.Println(fmt.Sprintf("[%v] %v coins %v total", blockNumPad, subPad, totalSubsidyPad))
		}
	}
	fmt.Println(fmt.Sprintf("totalSubsidy: %16v", totalSubsidy.ToFloat64()))
	fmt.Println("")
	return totalSubsidy
}

func calcSubsidy(subsidyBlocksNumber int64, height int64, totalSubsidy float64) BigNum {
	if height == 0 { //genesis block
		return NewBigNum(0)
	}
	H := NewBigNum(height - 1)
	N := NewBigNum(subsidyBlocksNumber)

	//lastBlockIndex := new(big.Int).SetInt64(subsidyBlocksNumber - 1)
	lastBlockIndex := NewBigNum(subsidyBlocksNumber - 1)
	if H.Cmp(lastBlockIndex) > 0 {
		return NewBigNum(0)
	}
	//endSubsidy := float64(0)               // 0 coins

	//return totalSubsidy * 2.0 * (lastBlockIndex - H) / (N * lastBlockIndex)
	H = H.Neg(H)                 // -H
	H = H.Add(lastBlockIndex, H) // (lastBlockIndex - H)
	H = H.Mul(NewBigNum(2), H)   // 2.0 * (lastBlockIndex - H)
	N = N.Mul(N, lastBlockIndex) // (N * lastBlockIndex)

	//subsidy := big.NewRat(1, 1)
	subsidy := NewBigNum(1)
	subsidy = subsidy.SetFrac(H, N) //  2.0 * (lastBlockIndex - H) / (N * lastBlockIndex)

	T := NewBigNum(totalSubsidy)
	//T = T.SetFloat64(totalSubsidy)
	T = T.Mul(T, subsidy) // totalSubsidy * 2.0 * (lastBlockIndex - H) / (N * lastBlockIndex)

	//float64Result, _ := T.Float64()
	return T
	//return totalSubsidy * 2.0 * (lastBlockIndex - H) / (N * lastBlockIndex)
}
