package main

import (
	"fmt"
	"math/big"

	"github.com/jackytck/projecteuler/tools"
)

func solve(digits int) int {
	m := tools.Exp(10, digits)
	ans := big.NewInt(2)
	ans.Exp(ans, big.NewInt(7830457), m)
	ans.Mul(ans, big.NewInt(28433))
	ans.Add(ans, big.NewInt(1))
	ans.Mod(ans, m)
	return int(ans.Int64())
}

func main() {
	fmt.Println(solve(10))
}

// Find the last ten digits of the prime number 28433×2^7830457+1.
// Note:
// Straightforward using math/big.
