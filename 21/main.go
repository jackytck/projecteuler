package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func sumAmicable(under int) int {
	var s int
	m := make(map[int]int)
	for i := 2; i < under; i++ {
		m[i] = tools.SumDivisors(i)
	}
	for k, v := range m {
		if k == v {
			continue
		}
		if v < under {
			if k == m[v] {
				s += k
			}
		} else {
			if k == tools.SumDivisors(v) {
				s += k
			}
		}
	}
	return s
}

func main() {
	fmt.Println(sumAmicable(10000))
}

// Sum of all amicable numbers under n.
