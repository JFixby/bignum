package bignum

type BigNum interface {
	Add(a BigNum, b BigNum) BigNum
	SetFrac(x BigNum, y BigNum) BigNum
	Mul(x BigNum, y BigNum) BigNum
	Neg(num BigNum) BigNum
	Cmp(num BigNum) int
	ToFloat64() float64
	ToInt64() int64
}

type BigNumEngine interface {
	NewBigNum(value interface{}) BigNum
}

