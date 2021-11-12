//大数比较
package bignum

import "math/big"

// >= 0
func IsNotNegative(s string) bool {
	f, ok := new(big.Float).SetString(s)
	if !ok {
		return false
	}

	zero, _ := new(big.Float).SetString("0")
	if f.Cmp(zero) < 0 {
		return false
	}

	return true
}

// <= 0
func IsNotPositive(s string) bool {
	f, ok := new(big.Float).SetString(s)
	if !ok {
		return false
	}

	zero, _ := new(big.Float).SetString("0")
	if f.Cmp(zero) > 0 {
		return false
	}

	return true
}

// >0
func IsPositive(s string) bool {
	f, ok := new(big.Float).SetString(s)
	if !ok {
		return false
	}

	zero, _ := new(big.Float).SetString("0")
	if f.Cmp(zero) > 0 {
		return true
	}

	return false
}

// <0
func IsNegative(s string) bool {
	f, ok := new(big.Float).SetString(s)
	if !ok {
		return false
	}

	zero, _ := new(big.Float).SetString("0")
	if f.Cmp(zero) < 0 {
		return true
	}

	return false
}
