package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func quadratic(a, b int) func(int) int {
	return func(n int) int {
		return n*n + a*n + b
	}
}

func primeLength(q func(int) int) int {
	var len int
	for {
		if tools.IsPrime(q(len)) {
			len++
		} else {
			break
		}
	}
	return len
}

func main() {
	var maxLen, prod int
	for a := -999; a < 1000; a++ {
		for b := -999; b < 1000; b++ {
			q := quadratic(a, b)
			if len := primeLength(q); len > maxLen {
				maxLen = len
				prod = a * b
			}
		}
	}
	fmt.Println(prod)
}

// Find the quadratic that produces the maximum number of primes.
