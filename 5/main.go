package main

import (
	"fmt"
	"math"
)

func primeFactors(n int) map[int]int {
	factors := make(map[int]int)
	d := 2
	for n > 1 {
		for n%d == 0 {
			n /= d
			factors[d]++
		}
		d++
	}
	return factors
}

// Smallest positive number that is divisible by all of the numbers from 1 to n.
func lcm(n int) int {
	lcm := make(map[int]int)
	for i := 2; i <= n; i++ {
		for k, v := range primeFactors(i) {
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
