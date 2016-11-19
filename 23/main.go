package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func isAbundant(n int) bool {
	return tools.SumDivisors(n) > n
}

func isSumOfTwoAbundant(n int, a []bool) bool {
	var ret bool
	for i := 1; i < n; i++ {
		if a[i] && a[n-i] {
			ret = true
			break
		}
	}
	return ret
}

func solve() int {
	var sum int
	ab := make([]bool, 28124)
	for i := 12; i <= 28123; i++ {
		ab[i] = isAbundant(i)
	}
	for i := 1; i <= 28123; i++ {
		if !isSumOfTwoAbundant(i, ab) {
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Println(solve())
}

// Sum of all +ive integers which cannot be written as the sum of two abundant
// numbers.
