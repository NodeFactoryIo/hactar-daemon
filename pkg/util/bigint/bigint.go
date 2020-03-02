package bigint

import (
	"fmt"
	"math/big"
)

var EmptyInt = BigInt{}

type BigInt struct {
	*big.Int
}

func NewInt(i uint64) BigInt {
	return BigInt{big.NewInt(0).SetUint64(i)}
}

func BigFromBytes(b []byte) BigInt {
	i := big.NewInt(0).SetBytes(b)
	return BigInt{i}
}

func BigFromString(s string) (BigInt, error) {
	v, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		return BigInt{}, fmt.Errorf("failed to parse string as a big int")
	}

	return BigInt{v}, nil
}

func BigMul(a, b BigInt) BigInt {
	return BigInt{big.NewInt(0).Mul(a.Int, b.Int)}
}

func BigDiv(a, b BigInt) BigInt {
	return BigInt{big.NewInt(0).Div(a.Int, b.Int)}
}

func BigMod(a, b BigInt) BigInt {
	return BigInt{big.NewInt(0).Mod(a.Int, b.Int)}
}

func BigAdd(a, b BigInt) BigInt {
	return BigInt{big.NewInt(0).Add(a.Int, b.Int)}
}

func BigSub(a, b BigInt) BigInt {
	return BigInt{big.NewInt(0).Sub(a.Int, b.Int)}
}

func BigCmp(a, b BigInt) int {
	return a.Int.Cmp(b.Int)
}

func (bi BigInt) Nil() bool {
	return bi.Int == nil
}

// LessThan returns true if bi < o
func (bi BigInt) LessThan(o BigInt) bool {
	return BigCmp(bi, o) < 0
}

// GreaterThan returns true if bi > o
func (bi BigInt) GreaterThan(o BigInt) bool {
	return BigCmp(bi, o) > 0
}

// Equals returns true if bi == o
func (bi BigInt) Equals(o BigInt) bool {
	return BigCmp(bi, o) == 0
}
