package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func sumPrimes(n int) uint64 {
	var sum uint64
	for _, v := range tools.SievePrime(n) {
		sum += uint64(v)
	}
	return sum
}

func main() {
	fmt.Println(sumPrimes(10))
	fmt.Println(sumPrimes(2000000))
}

// Sum of all primes below n.
