package main

import (
	"fmt"
	"math"

	"github.com/jackytck/projecteuler/tools"
)

func solve(limit int) (int, int) {
	// largest r/s smaller than a/b
	var r, s int
	a, b := 3, 7
	// loop all p/q
	for q := 2; q <= limit; q++ {
		f := float64(a*q-1) / float64(b)
		p := int(math.Floor(f))
		if r == 0 || p*s > q*r {
			r, s = p, q
		}
	}
	return tools.SimplifyFraction(r, s)
}

func main() {
	fmt.Println(solve(8))
	fmt.Println(solve(1000000))
}

// By listing the set of reduced proper fractions for d ≤ 1,000,000 in ascending
// order of size, find the numerator of the fraction immediately to the left of
// 3/7.
// Note:
// Instead of looping over p and q, one could loop over q only, then infer p.
