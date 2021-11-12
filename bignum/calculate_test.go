package bignum

import (
	"testing"
)

const (
	bigFloatPositive = "9999999999.999"
	bigFloatNegative = "-9999999999.999"
)

func TestFloatOpposite(t *testing.T) {
	r, isSuc := FloatOpposite(bigFloatPositive)
	if !isSuc {
		t.Fatalf("fail")
	}

	t.Log(r)
}

func TestFloatAbs(t *testing.T) {
	r, isSuc := FloatAbs(bigFloatNegative)
	if !isSuc {
		t.Fatalf("fail")
	}

	t.Log(r)
}

func TestNegative(t *testing.T) {
	r, isSuc := Negative(bigFloatPositive)
	if !isSuc {
		t.Fatalf("fail")
	}

	t.Log(r)
}

func TestNewBigFloat(t *testing.T) {
	r, err := NewBigFloat(bigFloatPositive)
	if err != nil {
		t.Fatalf("err: %+v\n", err)
	}
	t.Log(r.String())

	if err := r.Add(bigFloatNegative); err != nil {
		t.Fatalf("err: %+v\n", err)
	}
	t.Log(r.String())

	if err := r.Sub(bigFloatNegative); err != nil {
		t.Fatalf("err: %+v\n", err)
	}
	t.Log(r.String())

	if err := r.Mul(bigFloatPositive); err != nil {
		t.Fatalf("err: %+v\n", err)
	}
	t.Log(r.String())

	if err := r.Quo(bigFloatPositive); err != nil {
		t.Fatalf("err: %+v\n", err)
	}
	t.Log(r.String())

	res, err := Cmp(bigFloatNegative, bigFloatPositive)
	if err != nil {
		t.Fatalf("err: %+v\n", err)
	}
	t.Log(res)
}
