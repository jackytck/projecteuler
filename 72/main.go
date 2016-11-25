package main

import "fmt"

func solve(limit int) int64 {
	var cnt int64
	phi := make([]int, limit+1)
	for i := 2; i <= limit; i++ {
		phi[i] = i
	}
	for i, p := range phi {
		// check for prime
		if i >= 2 && i == p {
			// loop multiple of prime
			for j := i; j <= limit; j += p {
				// use Euler's product formula
				phi[j] = phi[j] * (p - 1) / p
			}
		}
		cnt += int64(phi[i])
	}
	return cnt
}

func main() {
	fmt.Println(solve(8))
	fmt.Println(solve(1000000))
}

// Consider the fraction, n/d. How many elements would be contained in the set
// of reduced proper fractions for d ≤ 1,000,000?
// Note:
// The sum of the number of reduced proper fraction is the same as the sum of
// the Euler's totient function φ(n). So we need to find all the prime factors
// of all numbers under consideration. To do this efficiently, the sieve of
// prime idea is used. Given a prime p, the multiple of p, e.g. 2p, 3p, 4p...
// have a prime factor of p. So the φ(kp) could be incrementally multiplied.
