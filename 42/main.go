package main

import (
	"fmt"
	"log"

	"github.com/jackytck/projecteuler/tools"
)

func solve(path string) int {
	var cnt int
	words, err := tools.ReadWords(path)
	if err != nil {
		log.Fatal(err)
	}
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
