package main

import (
	"fmt"
	"math"

	"github.com/jackytck/projecteuler/tools"
)

// Smallest positive number that is divisible by all of the numbers from 1 to n.
func lcm(n int) int {
	lcm := make(map[int]int)
	for i := 2; i <= n; i++ {
		for k, v := range tools.PrimeFactors(i) {
			if v > lcm[k] {
				lcm[k] = v
			}
		}
	}
	p := 1.0
	for k, v := range lcm {
		p *= math.Pow(float64(k), float64(v))
	}
	return int(p)
}

func main() {
	fmt.Println(lcm(10))
	fmt.Println(lcm(20))
}

// What is the smallest positive number that is evenly divisible by all of the
// numbers from 1 to 20?
