package main

import (
	"fmt"
	"math"

	"github.com/jackytck/projecteuler/tools"
)

func main() {
	limit := 10000
	primes := tools.SievePrime(limit)
	m := make(map[int]bool)
	for _, p := range primes {
		m[p] = true
	}
	for i := 3; i < limit; i += 2 {
		if _, isPrime := m[i]; isPrime {
			continue
		}
		conjecture := false
		for _, p := range primes {
			if p >= i {
				break
			}
			if t := math.Sqrt((float64(i - p)) / 2); t == math.Floor(t) {
				conjecture = true
				// fmt.Printf("%d = %d + 2 * %d^2\n", i, p, int(t))
				break
			}
		}
		if !conjecture {
			fmt.Println(i)
			break
		}
	}
}

// The smallest odd composite that cannot be written as the sum of a prime and
// twice a square.
// Note:
// This solution just uses brute force to try all possible primes for each odd composite.
