package bignum

type BigNum interface {
	Add(a BigNum, b BigNum) BigNum
	SetFrac(x BigNum, y BigNum) BigNum
	Mul(x BigNum, y BigNum) BigNum
	Neg(num BigNum) BigNum
	Cmp(num BigNum) int
}

type BigNumEngine interface {
	ONE() BigNum
	ZERO() BigNum
	TWO() BigNum
	NewBigNum(value interface{}) BigNum
}

var Engine = defaultBigNumEngine()

func defaultBigNumEngine() BigNumEngine {
	return &Float64Engine{}
}

func ZERO() BigNum {
	return Engine.ZERO()
}
func ONE() BigNum {
	return Engine.ONE()
}
func TWO() BigNum {
	return Engine.TWO()
}

func NewBigNum(value interface{}) BigNum {
	return Engine.NewBigNum(value)
}
