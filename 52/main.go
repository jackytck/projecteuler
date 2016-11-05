package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func check(x int) bool {
	ans := true
	for i := 2; i <= 6; i++ {
		if !tools.IsPermuted(x, i*x) {
			ans = false
			break
		}
	}
	return ans
}

func main() {
	i := 2
	for !check(i) {
		i++
	}
	fmt.Println(i, i*2, i*3, i*4, i*5, i*6)
}

// The smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x, contain
// the same digits.
