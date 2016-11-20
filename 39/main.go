package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func solve(limit int) int {
	var maxP int
	var maxAns int
	for p := 2; p <= limit; p += 2 {
		var ans int
		for a := 1; a < p/3; a++ {
			if b, r := tools.Divmod(p*(p-2*a), 2*(p-a)); r == 0 && b >= a {
				ans++
			}
		}
		if ans > maxAns {
			maxAns = ans
			maxP = p
		}
	}
	return maxP
}

func main() {
	fmt.Println(solve(1000))
}

// Find the perimeter p of a right angle triangle, so that it gives the
// maximised number of integer solutions.
// Note:
// p must be even. And a <= b < c => a < p/3.
