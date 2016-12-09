package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

type triple struct {
	a, b, c int
}

func pythagorean(bound int) chan triple {
	yield := make(chan triple)
	go func() {
		for m := 1; m < bound; m++ {
			for n := 1; n < m; n++ {
				// m and n are coprime and not both odd
				if (m%2 == 1 && n%2 == 1) || tools.GCD(m, n) != 1 {
					continue
				}
				a := m*m - n*n
				b := 2 * m * n
				c := m*m + n*n
				if a > b {
					a, b = b, a
				}
				yield <- triple{a, b, c}
			}
		}
		close(yield)
	}()
	return yield
}

func solve(target, limit int) int {
	m := make([]int, limit)
	for t := range pythagorean(limit / 2) {
		// generate all triples
		for k := 1; true; k++ {
			tk := triple{t.a * k, t.b * k, t.c * k}
			if tk.b >= limit {
				break
			}
			// M = a
			if 2*tk.a >= tk.b {
				m[tk.a] += tk.a - (tk.b+1)/2 + 1
			}
			// M = b
			m[tk.b] += tk.a / 2
		}
	}
	var cnt int
	for i, v := range m {
		cnt += v
		if cnt >= target {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(solve(2000, 200))
	fmt.Println(solve(1000000, 5000))
}

// It can be shown that there are exactly 2060 distinct cuboids, ignoring
// rotations, with integer dimensions, up to a maximum size of M by M by M, for
// which the shortest route has integer length when M = 100. This is the least
// value of M for which the number of solutions first exceeds two thousand; the
// number of solutions when M = 99 is 1975.
// Find the least value of M such that the number of solutions first exceeds one
// million.
// Note:
// By flatterning the cuboid, the shortest path is actually the hypotenuse of a
// right-angled triangle. Assume the answer is smaller than 5000. Generate all
// the pythagoreans below that. And count how many solutions could be found for
// each pythagorean.
