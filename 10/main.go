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

func sumPrimes(n int) uint64 {
	var sum uint64
	for _, v := range sieve(n) {
		sum += uint64(v)
	}
	return sum
}

func main() {
	fmt.Println(sumPrimes(10))
	fmt.Println(sumPrimes(2000000))
}

// Sum of all primes below n.
