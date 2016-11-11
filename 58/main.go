package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func main() {
	side := 3
	prime := 3
	corner := 9
	for 10*prime > 2*side-1 {
		side += 2
		step := side - 1
		for i := 0; i < 3; i++ {
			corner += step
			if tools.IsPrime(corner) {
				prime++
			}
		}
		corner += step
	}
	fmt.Println(side)
}

// What is the side length of the square spiral for which the ratio of primes
// along both diagonals first falls below 10%?
