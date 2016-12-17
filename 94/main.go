package main

import "fmt"

func solve(limit int64) int64 {
	var x, y, cnt int64 = 2, 1, 0
	for 2*x-1 <= limit {
		// check if side and area are integers
		// b = a + 1
		if s3, a3 := 2*x-1, (x-2)*y; s3 > 0 && a3 > 0 && s3%3 == 0 && a3%3 == 0 {
			cnt += s3 + 1
		}
		// b = a - 1
		if s3, a3 := 2*x+1, (x+2)*y; s3 > 0 && a3 > 0 && s3%3 == 0 && a3%3 == 0 {
			cnt += s3 - 1
		}
		x, y = 2*x+3*y, 2*y+x
	}
	return cnt
}

func main() {
	fmt.Println(solve(1000000000))
}

// Find the sum of the perimeters of all almost equilateral triangles with
// integral side lengths and area and whose perimeters do not exceed one billion.
// Note:
// The solutions for the almost equilateral triangle could be found by solving
// the Pell's equation in the space of substituted variables, then check if the
// original variables are integers.
