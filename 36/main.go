package main

import (
	"fmt"
	"strconv"

	"github.com/jackytck/projecteuler/tools"
)

func solve(limit int) int {
	var sum int
	for i := 1; i < limit; i++ {
		b := strconv.FormatInt(int64(i), 2)
		if tools.IsPalindromeInt(i) && tools.IsPalindromeString(b) {
			sum += i
		}
	}
	return sum
}
func main() {
	fmt.Println(solve(1000000))
}

// Sum of all numbers, less than one million, which are palindromic in base 10
// and base 2.
