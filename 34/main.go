package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func main() {
	var sum int
	for i := 10; i < 10000000; i++ {
		var s int
		for _, d := range tools.Digits(i) {
			s += int(tools.Factorial(d).Int64())
			if s > i {
				break
			}
		}
		if s == i {
			sum += s
		}
	}
	fmt.Println(sum)
}

// Sum of all numbers which are equal to the sum of the factorial of their digits.
// Note:
// It could not have more than 7 digits because 9! * 8  = 2903040 is a 7 digits number.
