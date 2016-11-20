package main

import (
	"fmt"
	"math/big"

	"github.com/jackytck/projecteuler/tools"
)

func solve(limit int) int {
	var maxD int
	// largest value of x among all minimal solutions
	maxX := big.NewInt(0)
	for d := 2; d <= limit; d++ {
		if tools.IsSquareNumber(d) {
			continue
		}
		// minimal solution of x and y
		minX, minY := big.NewInt(0), big.NewInt(0)
		for i := 1; true; i++ {
			x, y := tools.ConvergentSqrt(d, i)
			minX.Set(x)
			minY.Set(y)
			x.Mul(x, x)
			y.Mul(y, y)
			y.Mul(y, big.NewInt(int64(-d)))
			if x.Add(x, y); x.Int64() == 1 {
				break
			}
		}
		if maxD == 0 || minX.Cmp(maxX) == 1 {
			maxD, maxX = d, minX
		}
	}
	// fmt.Println(maxD, maxX, maxY)
	return maxD
}

func main() {
	fmt.Println(solve(1000))
}

// Consider quadratic Diophantine equations of the form:
// x^2 – D*y^2 = 1
// Find the value of D ≤ 1000 in minimal solutions of x for which the largest
// value of x is obtained.
