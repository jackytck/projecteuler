package main

import (
	"fmt"
	"math"

	"github.com/jackytck/projecteuler/tools"
)

func largestPalindrome(digit int) (int, int, int) {
	start := int(math.Pow10(digit) - 1)
	end := int(math.Pow10(digit - 1))
	max, mi, mj := 0, 0, 0
	for i := start; i >= end; i-- {
		for j := start; j >= end; j-- {
			p := i * j
			if tools.IsPalindromeInt(p) {
				if max == 0 || p > max {
					max = p
					mi = i
					mj = j
				}
			}
		}
	}
	return max, mi, mj
}

func main() {
	fmt.Println(largestPalindrome(2))
	fmt.Println(largestPalindrome(3))
}

// The largest palindrome from the product of two n-digit numbers.
