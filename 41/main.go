package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func pandigitalPrime(upper int) int {
	var ans int
	primes := tools.SievePrime(upper)
	for i := len(primes) - 1; i >= 0; i-- {
		p := primes[i]
		if tools.IsPandigital(p) {
			ans = p
			break
		}
	}
	return ans
}

func main() {
	fmt.Println(pandigitalPrime(7654321))
}

// Find the largest n-digit pandigital prime.
// Note:
// If it has 9-digit, then its digit sum is 45, which is divisible by 3, which is
// not a prime. So it must not have 9 digits. Same for 8-digit. But it may have
// 7 digits. So we start from there.
