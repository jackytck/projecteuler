package main

import (
	"fmt"
	"math/big"

	"github.com/jackytck/projecteuler/tools"
)

func convergent(n int) *big.Int {
	a := big.NewInt(2)
	b := big.NewInt(3)
	if n == 1 {
		return a
	}
	if n == 2 {
		return b
	}
	for i := 3; i <= n; i++ {
		k := 1
		if i%3 == 0 {
			k = 2 * i / 3
		}
		k2 := big.NewInt(int64(k))
		c := big.NewInt(1).Set(b)
		c.Mul(c, k2)
		c.Add(c, a)
		a, b = b, c
	}
	return b
}

func solve(nth int) int {
	z := convergent(nth)
	// fmt.Println(z)
	// fmt.Println(tools.DigitSumBig(z))
	return int(tools.DigitSumBig(z).Int64())
}

func main() {
	fmt.Println(solve(100))
}

// Find the sum of digits in the numerator of the 100th convergent of the
// continued fraction for e.
// Note:
// The ith convergent equals p_i/q_i, where p_i and q_i can be calculated
// recursively using:
// p_n = a_n*p_{n−1} + p_{n−2}, q_n = a_n*q_{n−1} + q_{n−2}
