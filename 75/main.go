package main

import (
	"fmt"
	"math"

	"github.com/jackytck/projecteuler/tools"
)

func count(limit int) int {
	visited := make([]int, limit+1)
	mBound := int(math.Sqrt(float64(limit) / 2))
	for m := 1; m < mBound; m++ {
		for n := 1; n < m; n++ {
			if (m%2 == 1 && n%2 == 1) || tools.GCD(m, n) != 1 {
				continue
			}
			a := m*m - n*n
			b := 2 * m * n
			c := m*m + n*n
			p := a + b + c
			if p > limit {
				continue
			}
			pk := p
			for pk <= limit {
				visited[pk]++
				pk += p
			}
		}
	}
	var cnt int
	for _, v := range visited {
		if v == 1 {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(count(1500000))
}

// Given that L is the length of the wire, for how many values of L ≤ 1,500,000
// can exactly one integer sided right angle triangle be formed?
// Note:
// Use Euclid's formula to generate all the primitive triples:
// https://en.wikipedia.org/wiki/Pythagorean_triple#Generating_a_triple
// Then get all the triples from the mutliples of the primitives.
