package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func solve() string {
	var ans string
	var p4 []int
	isPrime := make(map[int]bool)
	for _, p := range tools.SievePrime(10000) {
		if p > 1000 {
			p4 = append(p4, p)
			isPrime[p] = true
		}
	}
	for i := 0; i < len(p4)-1; i++ {
		for j := i + 1; j < len(p4); j++ {
			a := p4[i]
			b := p4[j]
			c := b - a + b
			if isPrime[c] && tools.IsPermute(a, b, c) && a != 1487 {
				ans = fmt.Sprintf("%d%d%d", a, b, c)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(solve())
}

// Find three 4-digit primes that are evenly spaced and are permutation of each
// other.
