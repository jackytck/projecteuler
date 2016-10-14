package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func main() {
	low := 2
	up := 354294 // 9**5 * 6
	var sum int
	for i := low; i <= up; i++ {
		if int(tools.DigitSum(i, 5).Int64()) == i {
			sum += i
		}
	}
	fmt.Println(sum)
}

// Sum of all the numbers that can be written as the sum of fifth powers of their digits.
// Note:
// If the desired number has seven digits, but 9**5 * 7 = 413343 has only 6 digits.
// So it could not have seven digits. Same for eight, nine, ... digits of number.
// So a loose upper searching bound is 9**5 * 6 = 354294.
