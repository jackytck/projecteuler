package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func divisorsSquared(n int) int {
	c := 1
	p := tools.PrimeFactors(n)
	for _, v := range p {
		c *= 2*v + 1
	}
	return c
}

func solve(target int) int {
	i := target
	for divisorsSquared(i) < 2*target {
		i++
	}
	return i
}

func main() {
	fmt.Println(solve(2))
	fmt.Println(solve(1000))
}

// In the following equation x, y, and n are positive integers.
// 1/x + 1/y = 1/n
// What is the least value of n for which the number of distinct solutions
// exceeds one-thousand?
// Note:
// This is equivalent to asking the least n, where the number of divisors of n^2
// is larger than 2*target.
