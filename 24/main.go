package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func solve(nth int) int {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var i int
	for v := range tools.Perms(a) {
		i++
		if i == nth {
			return tools.JoinInts(v)
		}
	}
	return 0
}

func main() {
	fmt.Println(solve(1000000))
}

// The millionth lexicographic permutation of the digits: [0, 9].
