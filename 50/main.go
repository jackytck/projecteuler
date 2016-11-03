package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func consecutivePrimeSum(limit int) int {
	primes := tools.SievePrime(limit)
	m := make(map[int]bool)
	for _, p := range primes {
		m[p] = true
	}
	var ans int
	var longest int
	for i := 0; i < len(primes)-1; i++ {
		for j := i + 1; j < len(primes); j++ {
			s := tools.Sum(primes[i:j]...)
			if s > limit {
				break
			}
			if m[s] && j-i > longest {
				longest = j - i
				ans = s
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(consecutivePrimeSum(100))
	fmt.Println(consecutivePrimeSum(1000))
	fmt.Println(consecutivePrimeSum(1000000))
}

// Find the prime that has the longest sum of consecutive primes.
