package bignum

import (
	"fmt"
	"reflect"
)

type Float64Engine struct {
}

func (e Float64Engine) NewBigNum(v interface{}) BigNum {
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
	return &Float64Num{
		value: float64(floatVal),
	}
}

type Float64Num struct {
	value float64
}

func (n *Float64Num) Add(a BigNum, b BigNum) BigNum {
	af := a.(*Float64Num)
	bf := b.(*Float64Num)

	n.value = af.value + bf.value
	return n
}

func (n *Float64Num) SetFrac(a BigNum, b BigNum) BigNum {
	af := a.(*Float64Num)
	bf := b.(*Float64Num)

	n.value = af.value / bf.value
	return n
}

func (n *Float64Num) Mul(a BigNum, b BigNum) BigNum {
	af := a.(*Float64Num)
	bf := b.(*Float64Num)

	n.value = af.value * bf.value
	return n
}

func (n *Float64Num) Neg(num BigNum) BigNum {
	numf := num.(*Float64Num)
	n.value = -numf.value
	return n
}

func (n *Float64Num) Cmp(num BigNum) int {
	numf := num.(*Float64Num)
	if n.value > numf.value {
		return 1
	}
	if n.value < numf.value {
		return -1
	}
	return 0
}

func (n *Float64Num) String() string {
	return fmt.Sprintf("%v", n.value)
}
