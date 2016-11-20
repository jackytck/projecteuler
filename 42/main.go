package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/jackytck/projecteuler/tools"
)

func solve(path string) int {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var cnt int
	words := strings.Split(string(data), ",")
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
