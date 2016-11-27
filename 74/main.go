package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func facDigit(d int) int {
	if d < 0 || d > 9 {
		return 0
	}
	f := [10]int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
	return f[d]
}

func facSum(n int) int {
	var sum int
	for _, d := range tools.Digits(n) {
		sum += facDigit(d)
	}
	return sum
}

func chainLen(n int) int {
	var len int
	var last int
	for n != 169 && n != 363601 && n != 1454 && n != 871 && n != 45361 && n != 872 && n != 45362 && n != last {
		last, n = n, facSum(n)
		len++
	}
	if n != last {
		len += 2
	}
	if n == 169 || n == 363601 || n == 1454 {
		len++
	}
	return len
}

func count(limit, size int) int {
	var cnt int
	for i := 1; i < limit; i++ {
		if chainLen(i) == size {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(count(1000000, 60))
}

// How many chains, with a starting number below one million, contain exactly
// sixty non-repeating terms?
// Note:
// Any starting number will end up to 169, 363601, 1454, 871, 45361, 872, 45362
// or itself.
