package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func solve(limit int) int {
	set := make(map[int]bool)
	primes := tools.SievePrime(tools.SqrtInt(limit))
	for _, a := range primes {
		u := a * a
		for _, b := range primes {
			v := u + b*b*b
			if v >= limit {
				break
			}
			for _, c := range primes {
				w := v + c*c*c*c
				if w >= limit {
					break
				}
				set[w] = true
			}
		}
	}
	return len(set)
}

func main() {
	fmt.Println(solve(50))
	fmt.Println(solve(50000000))
}

// How many numbers below fifty million can be expressed as the sum of a prime
// square, prime cube, and prime fourth power?
// Note:
// Just brute force.
