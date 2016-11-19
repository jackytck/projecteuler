package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func nthPrime(n int) int {
	k := 2
	for {
		primes := tools.SievePrime(n * k)
		if len(primes) >= n {
			return primes[n-1]
		}
		k *= 2
	}
}

func main() {
	fmt.Println(nthPrime(6))
	fmt.Println(nthPrime(10001))
}

// The nth prime.
