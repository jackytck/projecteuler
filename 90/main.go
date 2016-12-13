package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func extend(d []int) []int {
	ret := d
	if tools.Includes(d, 6) && !tools.Includes(d, 9) {
		ret = append(d, 9)
	}
	if tools.Includes(d, 9) && !tools.Includes(d, 6) {
		ret = append(d, 6)
	}
	return ret
}

func isValidPair(d1, d2 []int, a, b int) bool {
	return tools.Includes(d1, a) && tools.Includes(d2, b) ||
		tools.Includes(d1, b) && tools.Includes(d2, a)
}

func isValid(d1, d2 []int) bool {
	d1, d2 = extend(d1), extend(d2)
	cases := []struct {
		a, b int
	}{
		{0, 1},
		{0, 4},
		{0, 9},
		{1, 6},
		{2, 5},
		{3, 6},
		{4, 9},
		{6, 4},
		{8, 1},
	}
	valid := true
	for _, c := range cases {
		if !isValidPair(d1, d2, c.a, c.b) {
			valid = false
			break
		}
	}
	return valid
}

func hash(d1, d2 []int) string {
	s1 := tools.JoinIntsString(d1...)
	s2 := tools.JoinIntsString(d2...)
	if s1 > s2 {
		return s2 + s1
	}
	return s1 + s2
}

func solve() int {
	set := make(map[string]bool)
	for d1 := range tools.CombIndex(10, 6) {
		for d2 := range tools.CombIndex(10, 6) {
			if isValid(d1, d2) {
				h := hash(d1, d2)
				set[h] = true
			}
		}
	}
	return len(set)
}

func main() {
	fmt.Println(solve())
}

// How many distinct arrangements of the two cubes allow for all of the square
// numbers to be displayed?
// Note:
// Space is small. Just brute force.
