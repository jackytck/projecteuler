package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func solve(path string) int {
	var cnt int
	words := tools.ReadWords(path)
	for _, w := range words {
		v := tools.WordValue(w)
		if tools.IsTriangleNumber(v) {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(solve("p042_words.txt"))
}

// Find the number of triangle words.
