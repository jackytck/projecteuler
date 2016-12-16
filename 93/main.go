package main

import (
	"fmt"
	"math"

	"github.com/jackytck/projecteuler/tools"
)

func add(a, b float64) float64 {
	return a + b
}

func minus(a, b float64) float64 {
	return a - b
}

func mul(a, b float64) float64 {
	return a * b
}

func div(a, b float64) float64 {
	if b == 0 {
		return math.MaxInt64
	}
	return a / b
}

func eval(a, b, c, d float64,
	op1 func(float64, float64) float64,
	op2 func(float64, float64) float64,
	op3 func(float64, float64) float64) []float64 {
	r1 := op1(a, op2(b, op3(c, d)))
	r2 := op1(a, op3(op2(b, c), d))
	r3 := op2(op1(a, b), op3(c, d))
	r4 := op3(op2(op1(a, b), c), d)
	r5 := op3(op1(a, op2(b, c)), d)
	return []float64{r1, r2, r3, r4, r5}
}

func count(m map[int]bool) int {
	cnt := 1
	for m[cnt] {
		cnt++
	}
	return cnt - 1
}

func solve() int {
	var maxSet []int
	var maxCnt int
	ops := []func(float64, float64) float64{add, minus, mul, div}
	// 10 choose 4 digits
	for digits := range tools.CombIndex(10, 4) {
		values := make(map[int]bool)
		// permutation of these 4 digits
		for p := range tools.Perms(digits) {
			// 3 Cartesian Product of 4 operators
			for op := range tools.CartProduct(4, 3) {
				a, b, c, d := float64(p[0]), float64(p[1]), float64(p[2]), float64(p[3])
				op1, op2, op3 := ops[op[0]], ops[op[1]], ops[op[2]]
				// 5 possible values from different parentheses arrangements
				for _, e := range eval(a, b, c, d, op1, op2, op3) {
					// total = 10C4 * 4! * 4^3 * 5 = 1612800
					if e > 0 && e == math.Floor(e) {
						values[int(e)] = true
					}
				}
			}
		}
		if cnt := count(values); maxCnt == 0 || cnt > maxCnt {
			maxCnt, maxSet = cnt, digits
		}
	}
	return tools.JoinInts(maxSet)
}

func main() {
	fmt.Println(solve())
}

// Find the set of four distinct digits, a < b < c < d, for which the longest
// set of consecutive positive integers, 1 to n, can be obtained, giving your
// answer as a string: abcd.
// Note:
// Since space is not large, just brute force all digits, operators and
// parentheses to get all possilbe values. Then find the largest consecutive set.
