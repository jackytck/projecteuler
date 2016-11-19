package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func countFactors(n int) int {
	count := 1
	p := tools.PrimeFactors(tools.TriangleNumber(n))
	for _, v := range p {
		count *= v + 1
	}
	return count
}

func search(over int) int {
	q := 1
	for countFactors(q) <= over {
		q++
	}
	return tools.TriangleNumber(q)
}

func main() {
	fmt.Println(search(5))
	fmt.Println(search(500))
}

// First triangle number to have over n divisors.
