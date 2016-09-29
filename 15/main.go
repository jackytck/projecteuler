package main

import (
	"fmt"
	"math/big"
)

func nCr(n, r int64) *big.Int {
	z := big.NewInt(0)
	return z.Binomial(n, r)
}

func main() {
	fmt.Println(nCr(4, 2))
	fmt.Println(nCr(40, 20))
}

// The number of ways of going from the top-left corner to the lower-right corner of a n-by-n grid.

// Note:
// Each step only has 2 variations, either right or down.
// For each path, the sum of each variation must be equal to n.
// But there is a total of 2n steps for each path.
// So the total number of paths is nCr(2n, n).
