package main

import (
	"fmt"
	"strconv"

	"github.com/jackytck/projecteuler/tools"
)

func main() {
	var sum int
	for i := 1; i < 1000000; i++ {
		b := strconv.FormatInt(int64(i), 2)
		if tools.IsPalindromeInt(i) && tools.IsPalindromeString(b) {
			sum += i
		}
	}
	fmt.Println(sum)
}

// Sum of all numbers, less than one million, which are palindromic in base 10 and base 2.
