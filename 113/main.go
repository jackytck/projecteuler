package main

import (
	"fmt"
	"math/big"

	"github.com/jackytck/projecteuler/tools"
)

func solve(n int) *big.Int {
	a := tools.NCR(n+10, 10)
	b := tools.NCR(n+9, 9)
	a.Add(a, b)
	a.Sub(a, big.NewInt(2))
	a.Sub(a, big.NewInt(int64(10*n)))
	return a
}

func main() {
	fmt.Println(solve(6))
	fmt.Println(solve(10))
	fmt.Println(solve(100))
}

// How many numbers below a googol (10^100) are not bouncy?
// Note:
// An application of lattice paths.
// http://mathworld.wolfram.com/LatticePath.html
