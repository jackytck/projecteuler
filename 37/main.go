package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
	lane "gopkg.in/oleiade/lane.v1"
)

func main() {
	// find right-truncatable primes first
	add := [4]int{1, 3, 7, 9}
	var right []int
	q := lane.NewDeque()
	q.Append(2)
	q.Append(3)
	q.Append(5)
	q.Append(7)
	for !q.Empty() {
		f := q.Shift().(int)
		for _, a := range add {
			t := f*10 + a
			if tools.IsPrime(t) {
				right = append(right, t)
				q.Append(t)
			}
		}
	}
	// then find left-truncatable primes
	var both []int
	for _, r := range right {
		d := tools.Digits(r)
		left := true
		for i := 1; i < len(d); i++ {
			a := tools.JoinInts(d[i:])
			if !tools.IsPrime(a) {
				left = false
				break
			}
		}
		if left {
			both = append(both, r)
		}
	}
	fmt.Println(tools.Sum(both...))
}

// Sum of all primes that are both left-truncatable and right-truncatable.
