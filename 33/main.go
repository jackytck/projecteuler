package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func solve() (int, int) {
	var numerators []int
	var denominators []int
	for i := 11; i < 100; i++ {
		for j := 10; j < i; j++ {
			// j/i == ab/cd
			a, b := tools.Divmod(j, 10)
			c, d := tools.Divmod(i, 10)
			//trivial
			if b == 0 && d == 0 {
				continue
			}
			// fmt.Println(a, b, c, d)
			if a == c && b*i == d*j {
				// ab/cd == b/d
				numerators = append(numerators, b)
				denominators = append(denominators, d)
			} else if a == d && b*i == c*j {
				// ab/cd == b/c
				numerators = append(numerators, b)
				denominators = append(denominators, c)
			} else if b == c && a*i == d*j {
				// ab/cd == a/d
				numerators = append(numerators, a)
				denominators = append(denominators, d)
			} else if b == d && a*i == c*j {
				// ab/cd == a/c
				numerators = append(numerators, a)
				denominators = append(denominators, c)
			}
		}
	}
	n := tools.ProdInts(numerators)
	d := tools.ProdInts(denominators)
	return tools.SimplifyFraction(n, d)
}

func main() {
	fmt.Println(solve())
}

// Find the non-trivial 2-digits cancelling fractions, multiply and simplify
// them, and find the denominator.
