package main

import (
	"fmt"
	"math"

	"github.com/jackytck/projecteuler/tools"
)

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func un(n int) int {
	var r int
	for i := 0; i <= 10; i++ {
		s := 1
		if i%2 == 1 {
			s = -1
		}
		r += s * pow(n, i)
	}
	return r
}

func solve() int {
	x := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	y := make([]float64, 11)
	for i := 0; i <= 10; i++ {
		y[i] = float64(un(i))
	}
	var bop float64
	for i := 2; i <= 11; i++ {
		p := tools.LagrangePoly(x[1:i], y[1:i])
		bop += p(float64(i))
	}
	return int(math.Floor(bop))
}

func main() {
	fmt.Println(solve())
}

// Consider the following tenth degree polynomial generating function:
// u_n = 1 − n + n^2 − n^3 + n^4 − n^5 + n^6 − n^7 + n^8 − n^9 + n^10
// Find the sum of FITs for the BOPs.
// Note:
// Lagrange interpolating polynomials are used to interpolate the points with
// increasing degrees.
