package main

import (
	"fmt"
	"math/big"

	"github.com/jackytck/projecteuler/tools"
)

func count(n, s int) int {
	a := tools.NCR(n, s)
	b := tools.NCR(n-s, s)
	c := tools.NCR(n, 2*s)
	d := tools.Catalan(int64(s))
	a.Mul(a, b)
	a.Quo(a, big.NewInt(2))
	d.Mul(d, c)
	a.Sub(a, d)
	return int(a.Int64())
}

func solve(n int) int {
	var cnt int
	for s := 2; s <= n/2; s++ {
		cnt += count(n, s)
	}
	return cnt
}

func main() {
	fmt.Println(solve(4))
	fmt.Println(solve(7))
	fmt.Println(solve(12))
}

// Assume that a given set contains n strictly increasing elements and it
// already satisfies the second rule. For n = 12, how many of the 261625 subset
// pairs that can be obtained need to be tested for equality?
// Note:
// This problem is actually asking the number of pair of subsets of the same
// size, where swapping may occur. This reminds me of the proof of Quicksort,
// where Catalan number is used.
