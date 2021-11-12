package bignum

import (
	"errors"
	"fmt"
	"math/big"
)

func FloatOpposite(src string) (string, bool) {
	f, ok := new(big.Float).SetString(src)
	if !ok {
		return "0", false
	}

	z := new(big.Float)
	z.Neg(f)

	return fmt.Sprintf("%v", z), true
}

func FloatAbs(src string) (string, bool) {
	f, ok := new(big.Float).SetString(src)
	if !ok {
		return "0", false
	}

	z := new(big.Float)
	z.Abs(f)

	return fmt.Sprintf("%v", z), true
}

func Negative(src string) (string, bool) {
	dest, ok := FloatAbs(src)
	if ok == false {
		return "", false
	}

	dest, ok = FloatOpposite(dest)
	if ok == false {
		return "", false
	}

	return dest, true
}

//////////////////////////////////////////////////////////

type BigFloat struct {
	Value *big.Float
}

//init
func NewBigFloat(src string) (*BigFloat, error) {
	f, ok := new(big.Float).SetString(src)
	if !ok {
		return nil, errors.New("invalid")
	}

	return &BigFloat{Value: f}, nil
}

//加法
func (f *BigFloat) Add(src string) error {
	s, ok := new(big.Float).SetString(src)
	if !ok {
		return errors.New("invalid")
	}

	f.Value = f.Value.Add(f.Value, s)
	return nil
}

//减法
func (f *BigFloat) Sub(src string) error {
	s, ok := new(big.Float).SetString(src)
	if !ok {
		return errors.New("invalid")
	}

	f.Value = f.Value.Sub(f.Value, s)
	return nil
}

//乘法
func (f *BigFloat) Mul(src string) error {
	s, ok := new(big.Float).SetString(src)
	if !ok {
		return errors.New("invalid")
	}

	f.Value = f.Value.Mul(f.Value, s)
	return nil
}

//除法
func (f *BigFloat) Quo(src string) error {
	s, ok := new(big.Float).SetString(src)
	if !ok {
		return errors.New("invalid")
	}

	f.Value = f.Value.Quo(f.Value, s)
	return nil
}

//比较
// -1 if left < right
//  0 if left == right
// +1 if left >  right
func Cmp(left, right string) (int, error) {
	l, ok := new(big.Float).SetString(left)
	if !ok {
		return 0, errors.New("invalid")
	}

	r, ok := new(big.Float).SetString(right)
	if !ok {
		return 0, errors.New("invalid")
	}
	////   -1 if x <  y
	////    0 if x == y (incl. -0 == 0, -Inf == -Inf, and +Inf == +Inf)
	////   +1 if x >  y
	////
	//func (x *Float) Cmp(y *Float) int {
	//

	return l.Cmp(r), nil
}

//string
func (f *BigFloat) String() string {
	return fmt.Sprintf("%v", f.Value)
}
