package main

import (
	"fmt"
	"math"

	"github.com/jackytck/projecteuler/tools"
)

func firstDigs(nth, digit int) int {
	const lp float64 = 0.20898764024997873
	const lsqrt5 float64 = 0.3494850021680094
	t := float64(nth)*lp - lsqrt5
	t = t - math.Floor(t) + float64(digit) - 1
	return int(math.Pow(10, t))
}

func solve() int {
	m, n := 1000000000, 100000000
	a, b := 1, 1
	i := 1
	for {
		a, b = b, (a+b)%m
		if a > n && tools.IsPandigital(a) && tools.IsPandigital(firstDigs(i+1, 9)) {
			break
		}
		i++
	}
	return i + 1
}

func main() {
	fmt.Println(solve())
}

// Given that F_k is the first Fibonacci number for which the first nine digits
// AND the last nine digits are 1-9 pandigital, find k.
// Note:
// First find the k where F_k (mod 10^9) is pandigital, then use the Fibonacci
// formula to find the approximated F_k, where the first 9 digits could be
// extracted from log arithmetic.
