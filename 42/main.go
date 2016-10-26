package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/jackytck/projecteuler/tools"
)

func main() {
	data, err := ioutil.ReadFile("p042_words.txt")
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
	fmt.Println(cnt)
}

// Find the number of triangle words.
