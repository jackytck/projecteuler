package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func circularPrimes(n int) int {
	var cnt int
	primes := tools.SievePrime(n)
	isPrime := make(map[int]bool)
	for _, p := range primes {
		isPrime[p] = true
	}
	for _, p := range primes {
		d := tools.Digits(p)
		circular := true
		// rotation
		for i := 0; i < len(d); i++ {
			var r []int
			r = append(r, d[i:]...)
			r = append(r, d[:i]...)
			a := tools.JoinInts(r)
			_, ok := isPrime[a]
			if !ok {
				circular = false
				break
			}
		}
		if circular {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(circularPrimes(100))
	fmt.Println(circularPrimes(1000000))
}

// Number of circular primes below one million.
