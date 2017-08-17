package gohelpmath

import "math/big"

// NewBigIntFromBytes creates a BigInt and set it's bytes to the provided value
func NewBigIntFromBytes(buf []byte) (bigInt *big.Int) {
	bigInt = &big.Int{}
	bigInt.SetBytes(buf)
	return
}

// NewBigIntFromInt64 creates a BigInt and set it's value to the provided value
func NewBigIntFromInt64(value int64) (bigInt *big.Int) {
	bigInt = &big.Int{}
	bigInt.SetInt64(value)
	return
}
