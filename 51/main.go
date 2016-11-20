package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func count(p, r int) (int, []int) {
	var pos []int
	digits := tools.Digits(p)
	for i, d := range digits {
		if d == r && i != len(digits)-1 {
			pos = append(pos, i)
		}
	}
	return len(pos), pos
}

func replace(p, r int, pos []int) int {
	digits := tools.Digits(p)
	for _, p := range pos {
		digits[p] = r
	}
	return tools.JoinInts(digits)
}

func candidate(p, r int, primes map[int]bool) (bool, []int) {
	var can []int
	size, pos := count(p, r)
	if size != 3 {
		return false, can
	}
	for i := r; i <= 9; i++ {
		x := replace(p, i, pos)
		if primes[x] {
			can = append(can, x)
		}
	}
	return len(can) == 8, can
}

func check(p int, primes map[int]bool) bool {
	for r := 0; r < 3; r++ {
		if valid, _ := candidate(p, r, primes); valid {
			return true
		}
	}
	return false
}

func solve() int {
	var ans int
	primes := tools.SievePrime(1000000)
	m := make(map[int]bool)
	for _, p := range primes {
		m[p] = true
	}
	for _, p := range primes {
		if p < 10000 {
			continue
		}
		if check(p, m) {
			ans = p
			break
		}
	}
	return ans
}

func main() {
	fmt.Println(solve())
}

// Find the smallest prime which, by replacing part of the number with the same
// digit, is part of an eight prime value family.
// Note:
// Assume the answer has 5 or 6 digits, then the number of repeating digits
// must be 3.
