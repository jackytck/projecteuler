package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func solve(limit int) int {
	ans := 1
	for _, p := range tools.SievePrime(50) {
		t := ans * p
		if t >= limit {
			break
		}
		ans = t
	}
	return ans
}

func main() {
	fmt.Println(solve(1000000))
}

// Find the value of n ≤ 1,000,000 for which n/φ(n) is a maximum.
// Note:
// Euler's product formula is used:
// https://en.wikipedia.org/wiki/Euler%27s_totient_function#Euler.27s_product_formula
