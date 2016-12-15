package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func next(n int) int {
	var c int
	for _, d := range tools.Digits(n) {
		c += d * d
	}
	return c
}

func solve(bound int) int {
	var cnt int
	m1 := make(map[int]bool)
	m89 := make(map[int]bool)
	for i := 2; i < bound; i++ {
		t := i
		for t != 1 && t != 89 && !m1[t] && !m89[t] {
			t = next(t)
		}
		if t == 1 || m1[t] {
			m1[i] = true
			// fmt.Println(i, t, 1)
		} else {
			m89[i] = true
			cnt++
			// fmt.Println(i, t, 89)
		}
	}
	return cnt
}

func main() {
	fmt.Println(solve(10000000))
}

// A number chain is created by continuously adding the square of the digits in
// a number to form a new number until it has been seen before. How many
// starting numbers below ten million will arrive at 89?
// Note:
// Just brute force and cache all the results.
