package main

import (
	"fmt"
	"time"
)

func triangle(n int) int {
	return n * (n + 1) / 2
}

func primeFactors(n int) map[int]int {
	factors := make(map[int]int)
	d := 2
	for n > 1 {
		for n%d == 0 {
			n /= d
			factors[d]++
		}
		d++
	}
	return factors
}

func countFactors(n int) int {
	count := 1
	p := primeFactors(triangle(n))
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
	return triangle(q)
}

func main() {
	fmt.Println(search(5))
	start := time.Now()
	fmt.Println(search(500), time.Since(start))
}

// First triangle number to have over n divisors.
