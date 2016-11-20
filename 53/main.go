package main

import (
	"fmt"
	"math/big"

	"github.com/jackytck/projecteuler/tools"
)

func solve() int {
	var cnt int
	million := big.NewInt(1000000)
	for n := 1; n <= 100; n++ {
		for r := 1; r <= n; r++ {
			z := tools.NCR(n, r)
			z.Sub(z, million)
			if z.Sign() > 0 {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(solve())
}

// Number of values of nCr, where 1 ≤ n ≤ 100, that are greater than one-million.
