package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func solve(bound int) int64 {
	var sum int64
	for i := 1; i <= bound; i++ {
		if tools.IsSquareNumber(i) {
			continue
		}
		sum += tools.DigitSumBig(tools.Sqrt(i, bound)).Int64()
	}
	return sum
}

func main() {
	fmt.Println(solve(100))
}

// For the first one hundred natural numbers, find the total of the digital sums
// of the first one hundred decimal digits for all the irrational square roots.
// Note:
// Square roots by subtraction is used to compute the digits precisely:
// http://www.afjarvis.staff.shef.ac.uk/maths/jarvisspec02.pdf
