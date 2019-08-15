package bignum

type Float64Engine struct {
	zero *Float64Num
	one  *Float64Num
	two  *Float64Num
}

type Float64Num struct {
	value float64
}

func (e Float64Engine) ZERO() *Float64Num {
	if e.zero == nil {
		e.zero = e.NewBigNum(0)
	}
	return e.zero
}

func (e Float64Engine) ONE() *Float64Num {
	if e.one == nil {
		e.one = e.NewBigNum(1)
	}
	return e.one
}

func (e Float64Engine) TWO() *Float64Num {
	if e.two == nil {
		e.two = e.NewBigNum(0)
	}
	return e.two
}

func (e Float64Engine) NewBigNum(value interface{}) *Float64Num {
	return &Float64Num{
		value: float64(value),
	}
}
