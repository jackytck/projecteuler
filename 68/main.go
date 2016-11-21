package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

// 		0
// 	 	  \
// 			  1     3
// 		  /   \  /
// 		8       2
// 	/	  \ __ /
// 9		 6  4 -- 5
// 			  \
// 				 7

func solve() string {
	var max string
	for p := range tools.Perms([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		if p[0] != 10 && p[3] != 10 && p[5] != 10 && p[7] != 10 && p[9] != 10 {
			continue
		}
		if p[0] > p[3] || p[0] > p[5] || p[0] > p[7] || p[0] > p[9] {
			continue
		}
		a := p[0] + p[1] + p[2]
		b := p[3] + p[2] + p[4]
		c := p[5] + p[4] + p[6]
		d := p[7] + p[6] + p[8]
		e := p[9] + p[8] + p[1]
		if a == b && b == c && c == d && d == e {
			s := tools.JoinIntsString(p[0], p[1], p[2], p[3], p[2], p[4], p[5], p[4],
				p[6], p[7], p[6], p[8], p[9], p[8], p[1])
			if s > max {
				max = s
			}
		}
	}
	return max
}

func main() {
	fmt.Println(solve())
}

// What is the maximum 16-digit string for a "magic" 5-gon ring?
// Note:
// Since it is 16-digit, so number 10 must be sitting in the outer ring.
