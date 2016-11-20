package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func solve() int {
	i := 144
	for {
		h := tools.HexagonNumber(i)
		if tools.IsPentagonNumber(h) {
			break
		}
		i++
	}
	return tools.TriangleNumber(2*i - 1)
}

func main() {
	fmt.Println(solve())
}

// After 40755, the next triangle number that is also pentagonal and hexagonal.
// Note:
// The i-th hexagonal number is the (2*i-1)-th triangle number.
