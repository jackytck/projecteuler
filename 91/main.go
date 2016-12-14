package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func solve(size int) int {
	// all the right-angles sitting on the axis
	cnt := 3 * size * size
	for x := 1; x <= size; x++ {
		for y := 1; y <= size; y++ {
			f := tools.GCD(x, y)
			a := y * f / x
			b := (size - x) * f / y
			// by mirroring, so times two
			cnt += tools.MinInt(a, b) * 2
		}
	}
	return cnt
}

func main() {
	fmt.Println(solve(2))
	fmt.Println(solve(50))
}

// The points P (x1, y1) and Q (x2, y2) are plotted at integer co-ordinates
// and are joined to the origin, O(0,0), to form ΔOPQ.
// Given that 0 ≤ x1, y1, x2, y2 ≤ 50, how many right triangles can be formed?
// Note:
// All the right triangles could be partitioned into two sets. The first set
// contains all the right triangle whose right-angles touch either the x-axis,
// y-axis or the origin. The second set contains all of the rest whose right-
// angles are either on P or Q. But each right-angle P is a transpose of a
// right-angle Q. So we only need to count the right-angles on P, then times 2.
