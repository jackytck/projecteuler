package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func solve(goal int) int {
	bound := tools.SqrtInt(goal)
	var ans, min int
	for i := 1; i <= bound; i++ {
		for j := 1; j <= bound; j++ {
			if d := tools.AbsInt(goal - i*(i+1)*j*(j+1)/4); min == 0 || d < min {
				ans, min = i*j, d
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(solve(2000000))
}

// Although there exists no rectangular grid that contains exactly two million
// rectangles, find the area of the grid with the nearest solution.
// Note:
// Since any two distinct horizontal and two vertical lines could form a
// rectangle, and any rectangle could be deconstructed into these four lines.
// So it is complete and no more or less. So for a grid of X-by-Y, the total
// number of rectangles is {X+1}_C_{2} * {Y+1}_C_{2}.
