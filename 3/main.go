package main

import "fmt"

func largestPrime(n uint64) uint64 {
	var largest uint64 = 2
	for n > 1 {
		for n%largest == 0 {
			n /= largest
		}
		largest++
	}
	return largest - 1
}

func main() {
	fmt.Println(largestPrime(13195))
	fmt.Println(largestPrime(600851475143))
}

// Find the largest prime factor of a number.

// Note:
// Every natural number has a unique prime factorization.
// Starting from low to high, divide the number repeatedly and consume all the factors.
// Starting from 2, which is a prime. Then consider the next integer.
// At some point, one of the next integer is another factor.
// Assume this factor is a composit, then it will have a smaller prime factor,
// which should be consumed in the past, so this leads to contradiction.
// So it must be a prime.
