package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func isBouncy(n int) bool {
	ds := tools.Digits(n)
	var up, down bool
	for i := 1; i < len(ds); i++ {
		if ds[i] > ds[i-1] {
			up = true
		} else if ds[i] < ds[i-1] {
			down = true
		}
		if up && down {
			break
		}
	}
	return up && down
}

func solve(percent int) int {
	var bouncy int
	total := 99
	for total*percent > 100*bouncy {
		total++
		if isBouncy(total) {
			bouncy++
		}
	}
	// fmt.Println(bouncy, total, float64(bouncy)/float64(total))
	return total
}

func main() {
	fmt.Println(solve(50))
	fmt.Println(solve(90))
	fmt.Println(solve(99))
}

// Find the least number for which the proportion of bouncy numbers is exactly
// 99%.
// Note:
// Straightforward brute force.
