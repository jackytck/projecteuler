package main

import "fmt"

func solve(limit int) int {
	// total 62 scores
	var all []int
	// 21 double scores
	var double []int
	for i := 1; i <= 20; i++ {
		all = append(all, i, i*2, i*3)
		double = append(double, i*2)
	}
	all = append(all, 25, 50)
	double = append(double, 50)

	var cnt int

	// one dart, at most 21
	for _, d1 := range double {
		if d1 < limit {
			cnt++
		}
	}

	// two darts, at most 62 * 21 = 1302
	for _, d1 := range all {
		for _, d2 := range double {
			if d1+d2 < limit {
				cnt++
			}
		}
	}

	// three darts, at most (62 + 62C2) * 21 = 41013
	for i, d1 := range all {
		for j, d2 := range all {
			if j < i || d1+d2 >= limit {
				continue
			}
			for _, d3 := range double {
				if d1+d2+d3 < limit {
					cnt++
				}
			}
		}
	}

	return cnt
}

func main() {
	fmt.Println(solve(171))
	fmt.Println(solve(100))
}

// In the game of darts a player throws three darts at a target board which is
// split into twenty equal sized sections numbered one to twenty. How many
// distinct ways can a player checkout with a score less than 100?
// Note:
// Count in 3 cases, 1 dart, 2 darts and 3 darts. Pretty straightforward.
