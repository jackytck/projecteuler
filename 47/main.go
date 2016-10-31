package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func consecutive(n int) int {
	m := make(map[int]int)
	i := 2
	for {
		p := tools.PrimeFactors(i)
		m[i] = len(p)
		if i > n {
			valid := true
			for j := 0; j < n; j++ {
				if m[i-j] != n {
					valid = false
					break
				}
			}
			if valid {
				break
			}
		}
		i++
	}
	return i - n + 1
}

func main() {
	fmt.Println(consecutive(2))
	fmt.Println(consecutive(3))
	fmt.Println(consecutive(4))
}

// The first four consecutive integers to have four distinct prime factors.
