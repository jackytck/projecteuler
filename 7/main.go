package main

import (
	"fmt"
	"math"
)

// The sieve of Eratosthenes: primes below n
func sieve(n int) []int {
	var primes []int
	comp := make([]bool, n+2)
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		for j := i * i; j < n; j += i {
			comp[j] = true
		}
	}
	for i, v := range comp {
		if i > 1 && !v && i < n {
			primes = append(primes, i)
		}
	}
	return primes
}

func nthPrime(n int) int {
	k := 2
	for {
		primes := sieve(n * k)
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
