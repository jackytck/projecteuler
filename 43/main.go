package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func extract(slice []int, start, end int) int {
	return tools.JoinInts(slice[start:end])
}

func check(n, d int) bool {
	return n%d == 0
}

func solve() int {
	var sum int
	for v := range tools.Perms([]int{0, 1, 2, 3, 4, 6, 7, 8, 9}) {
		p2 := extract(v, 1, 4)
		p3 := extract(v, 2, 5)
		p7 := v[4]*100 + 50 + v[5]
		p11 := 500 + extract(v, 5, 7)
		p13 := extract(v, 5, 8)
		p17 := extract(v, 6, 9)
		if check(p2, 2) && check(p3, 3) && check(p7, 7) && check(p11, 11) && check(p13, 13) && check(p17, 17) {
			t := append(v[:5], append([]int{5}, v[5:]...)...)
			sum += tools.JoinInts(t)
		}
	}
	return sum
}

func main() {
	fmt.Println(solve())
}

// Sum of all 0 to 9 pandigital numbers, where its sub-digits are divisible by
// 2, 3, 5, 7, 11, 13 and 17 respectively.
