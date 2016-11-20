package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func powerDigitSum(a, b int) int {
	return tools.Sum(tools.DigitsBig(tools.Exp(a, b))...)
}

func solve() int {
	var max int
	for a := 2; a < 100; a++ {
		for b := 1; b < 100; b++ {
			if s := powerDigitSum(a, b); s > max {
				max = s
			}
		}
	}
	return max
}

func main() {
	fmt.Println(solve())
}

// The maximum digital sum of a^b, where a, b < 100.
