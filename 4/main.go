package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func isPalindrome(num int) bool {
	s := strconv.Itoa(num)
	return reverse(s) == s
}

func largestPalindrome(digit int) (int, int, int) {
	start := int(math.Pow10(digit) - 1)
	end := int(math.Pow10(digit - 1))
	max, mi, mj := 0, 0, 0
	for i := start; i >= end; i-- {
		for j := start; j >= end; j-- {
			p := i * j
			if isPalindrome(p) {
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
