package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func count(limit int) int {
	var cnt int
	for d := 2; d <= limit; d++ {
		for n := 1; n < d; n++ {
			if 3*n > d && 2*n < d && tools.GCD(n, d) == 1 {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(count(8))
	fmt.Println(count(12000))
}

// How many fractions lie between 1/3 and 1/2 in the sorted set of reduced
// proper fractions for d ≤ 12,000?
// Note:
// Just brute force.
// Faster way is to use Farey Sequences.
