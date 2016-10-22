package main

import (
	"fmt"
	"strconv"

	"github.com/jackytck/projecteuler/tools"
)

func catProd(x int, n int) int {
	var s string
	for i := 1; i <= n; i++ {
		s += strconv.Itoa(x * i)
	}
	p, _ := strconv.Atoi(s)
	return p
}

func main() {
	var largest int
	for i := 1; i < 10000; i++ {
		j := 1
		for {
			p := catProd(i, j)
			if p > 999999999 {
				break
			}
			if tools.IsPandigital(p) {
				largest = tools.MaxInt(largest, p)
			}
			j++
		}
	}

	fmt.Println(largest)
}

// Largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an integer.
