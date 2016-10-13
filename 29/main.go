package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func distinct(r1, r2 int) int {
	m := make(map[string]bool)
	for a := r1; a <= r2; a++ {
		for b := r1; b <= r2; b++ {
			p := tools.Exp(a, b)
			m[p.String()] = true
		}
	}
	return len(m)
}

func main() {
	fmt.Println(distinct(2, 5))
	fmt.Println(distinct(2, 100))
}

// Distinct terms gernerated by a^b for 2 <= a <= 100 and 2 <= b <= 100.
