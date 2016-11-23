package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func solve() int {
	var minN, minPhi int
	primes := tools.SievePrime(5000)
	for idx := range tools.CombIndex(len(primes), 2) {
		p1, p2 := primes[idx[0]], primes[idx[1]]
		n := p1 * p2
		if n > 10000000 {
			continue
		}
		phi := (p1 - 1) * (p2 - 1)
		if tools.IsPermuted(n, phi) && (minN == 0 || minN*phi > n*minPhi) {
			minN, minPhi = n, phi
		}
	}
	return minN
}

func main() {
	fmt.Println(solve())
}

// Find the value of n, 1 < n < 10^7, for which φ(n) is a permutation of n and
// the ratio n/φ(n) produces a minimum.
// Note:
// Assume the answser is a product of two primes.
