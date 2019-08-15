package bignum

type BigNum interface {
	Add(a BigNum, b BigNum) BigNum
	SetFrac(x BigNum, y BigNum) BigNum
	Mul(x BigNum, y BigNum) BigNum
	Neg(num BigNum) BigNum
	Cmp(num BigNum) int
	ToFloat64() float64
}

type BigNumEngine interface {
	NewBigNum(value interface{}) BigNum
}

var Engine = defaultBigNumEngine()

func defaultBigNumEngine() BigNumEngine {
	return &Float64Engine{}
}

func NewBigNum(value interface{}) BigNum {
	return Engine.NewBigNum(value)
}
