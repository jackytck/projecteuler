package main

import (
	"fmt"
	"strconv"

	"github.com/jackytck/projecteuler/tools"
)

// n: number of digits
// d: repeating digit
// r: how many times d is repeated
func genStr(n, d, r int) []string {
	if n == 0 {
		return []string{""}
	}
	if r == n {
		var w string
		for i := 0; i < n; i++ {
			w += strconv.Itoa(d)
		}
		return []string{w}
	}
	var ret []string
	for i := 0; i < 10; i++ {
		if i == d {
			for _, g := range genStr(n-1, d, r-1) {
				ret = append(ret, strconv.Itoa(i)+g)
			}
		} else {
			for _, g := range genStr(n-1, d, r) {
				ret = append(ret, strconv.Itoa(i)+g)
			}
		}
	}
	return ret
}

func genPrimes(n, d, r int) []int {
	var p []int
	for _, s := range genStr(n, d, r) {
		g, _ := strconv.Atoi(s)
		if string(s[0]) != "0" && tools.IsPrime(g) {
			p = append(p, g)
		}
	}
	return p
}

// M gives the maximum number of repeated digits for an n-digit prime where d is
// the repeated digit
func M(n, d int) int {
	m := n
	for len(genPrimes(n, d, m)) == 0 {
		m--
	}
	return m
}

// N gives the number of primes described in M.
func N(n, d int) int {
	return len(genPrimes(n, d, M(n, d)))
}

// S gives the sum of primes described in M.
func S(n, d int) int64 {
	var cnt int64
	for _, p := range genPrimes(n, d, M(n, d)) {
		cnt += int64(p)
	}
	return cnt
}

func solve(n int) int64 {
	var s int64
	for d := 0; d < 10; d++ {
		s += S(n, d)
		// fmt.Println(d, M(n, d), N(n, d), S(n, d))
	}
	return s
}

func main() {
	fmt.Println(solve(4))
	fmt.Println(solve(10))
}

// Find the sum of all S(10, d).
// Note:
// e.g. n=4 and d=2, generate 2222, 0222, 1222, etc, in decending order of
// repeated d, until a prime is found. Then M(n, d) is known, which is closed to
// n. So the searching space is small. So we could exhaust all other possible
// digits after fixing M(n, d) digits d.
