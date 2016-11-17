package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func period(s int) int {
	return len(tools.SqrtExapnd(s)) - 1
}

func main() {
	var odd int
	for n := 1; n <= 10000; n++ {
		if period(n)%2 == 1 {
			odd++
		}
	}
	fmt.Println(odd)
}

// How many continued fractions for N ≤ 10000 have an odd period?
// Note:
// Refer to
// https://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Continued_fraction_expansion
// for continued fraction expansion.
