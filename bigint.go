package bignum

type BigIntEngine struct {
}

func (e BigIntEngine) ZERO() BigNum {
	return nil
}

func (e BigIntEngine) ONE() BigNum {
	return nil
}

func (e BigIntEngine) TWO() BigNum {
	return nil
}

func (e BigIntEngine) NewBigNum(value interface{}) BigNum {
	return nil
}
