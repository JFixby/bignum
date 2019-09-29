package bignum

import (
	"fmt"
	"math/big"
	"reflect"
)

type BigDecimalEngine struct {
}


func (n BigDecimalEngine) String() string {
	return "bignum.BigDecimalEngine"
}

func (e BigDecimalEngine) NewBigNum(v interface{}) BigNum {
	floatVal := 0.0
	switch v.(type) {
	case float64:
		floatVal = v.(float64)
	case int:
		floatVal = float64(v.(int))
	case int64:
		floatVal = float64(v.(int64))
	default:
		panic(fmt.Sprintf("Unsupported type: %v(%v)", v, reflect.TypeOf(v)))
	}

	value := big.NewRat(1, 1)
	value.SetFloat64(floatVal)
	return &RatNum{
		value: value,
	}
}

type RatNum struct {
	value *big.Rat
}

func (n *RatNum) Add(a BigNum, b BigNum) BigNum {
	af := a.(*RatNum)
	bf := b.(*RatNum)
	n.value.Add(af.value, bf.value)
	return n
}

func (n *RatNum) SetFrac(a BigNum, b BigNum) BigNum {
	af := a.(*RatNum)
	bf := b.(*RatNum)

	tmp := big.NewRat(1, 1)
	//
	//pin.D("a", a)
	//pin.D("b", b)
	//pin.D("n", n)
	tmp.Inv(bf.value)
	//pin.D("tmp", tmp)
	tmp.Mul(af.value, tmp)
	//pin.D("tmp", tmp)
	n.value.Set(tmp)
	//pin.D("n", n)
	return n
}

func (n *RatNum) Mul(a BigNum, b BigNum) BigNum {
	af := a.(*RatNum)
	bf := b.(*RatNum)
	n.value.Mul(af.value, bf.value)
	return n
}

func (n *RatNum) Neg(num BigNum) BigNum {
	numf := num.(*RatNum)
	n.value.Neg(numf.value)
	return n
}

func (n *RatNum) Cmp(num BigNum) int {
	numf := num.(*RatNum)
	return n.value.Cmp(numf.value)
}

func (n RatNum) String() string {
	return fmt.Sprintf("%v", n.value)
}

func (n *RatNum) ToFloat64() float64 {
	f, _ := n.value.Float64()
	return f
}

func (n *RatNum) ToInt64() int64 {
	return int64(n.ToFloat64())
}
