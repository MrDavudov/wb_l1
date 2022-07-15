package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1 << 32) // 4294967296
	b := big.NewInt(1 << 62) // 4611686018427387904

	ab := new(big.Int).Mul(a, b)
	fmt.Printf("%v * %v = %v\n", a, b, ab) // 4294967296 * 4611686018427387904 = 19807040628566084398385987584

	abPlusB := new(big.Int).Add(ab, b)
	fmt.Printf("%v + %v = %v\n", ab, b, abPlusB) // 19807040628566084398385987584 + 4611686018427387904 = 19807040633177770416813375488

	aMinusAb := new(big.Int).Sub(a, ab)
	fmt.Printf("%v - %v = %v\n", a, ab, aMinusAb) // 4294967296 - 19807040628566084398385987584 = -19807040628566084394091020288

	div := new(big.Int).Div(abPlusB, b)
	fmt.Printf("%v / %v = %v\n", abPlusB, b, div) // 19807040633177770416813375488 / 4611686018427387904 = 4294967297
}