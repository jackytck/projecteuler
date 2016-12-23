package main

import "fmt"

func solve(limit int64) int64 {
	var b, t int64 = 15, 21
	for t <= limit {
		b, t = 3*b+2*t-2, 4*b+3*t-3
	}
	return b
}

func main() {
	fmt.Println(solve(1000000000000))
}

// A box contains two kinds of disc. An arrangement could be set so that exactly
// 50% chance of taking two blue discs at random.
// By finding the first arrangement to contain over 10^12 = 1,000,000,000,000
// discs in total, determine the number of blue discs that the box would contain.
// Note:
// This is a problem of solving a diophantine equation:
// https://en.wikipedia.org/wiki/Diophantine_equation
// It could be solved by:
// https://www.alpertron.com.ar/QUAD.HTM
