package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func dp(over int) int {
	primes := tools.SievePrime(over)
	table := make(map[int]int)
	table[0] = 1
	for _, p := range primes {
		for i := p; i <= over; i++ {
			table[i] += table[i-p]
		}
	}
	for k := 2; k <= over; k++ {
		if table[k] > over {
			return k
		}
	}
	return 0
}

func main() {
	fmt.Println(dp(5000))
}

// What is the first value which can be written as the sum of primes in over
// five thousand different ways?
// Note:
// This is also the same DP as problem 31 and 76. The only difference is: now we
// are considering the primes as coin 2, 3, 5, 7...
// It is assumed that the answer is well below both the highest prime coin and
// the desired number of ways.
