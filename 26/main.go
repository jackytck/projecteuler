package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func recurringCycle(a, b int) ([]int, int) {
	var digits []int
	var period int
	m := make(map[int]int)
	var i int
	p := a
	for {
		var q int
		m[p] = i
		p, q = tools.Divmod(p*10, b)
		digits = append(digits, p)
		if v, ok := m[q]; ok {
			period = i - v + 1
			break
		}
		if q == 0 {
			break
		}
		p = q
		i++
	}
	return digits, period
}

func main() {
	var maxInt, maxCycle int
	for i := 2; i <= 1000; i++ {
		_, c := recurringCycle(1, i)
		if c > maxCycle {
			maxInt = i
			maxCycle = c
		}
	}
	fmt.Println(maxInt)
}

// The number d for which 1/d has the longest recurring cycle in its decimal fraction part.
