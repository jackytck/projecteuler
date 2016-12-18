package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func sumDivisors(limit int) []int {
	s := make([]int, limit+1)
	for i := 1; i <= limit/2; i++ {
		for j := i * 2; j <= limit; j += i {
			s[j] += i
		}
	}
	return s
}

func solve(limit int) int {
	sumDiv := sumDivisors(limit)
	var ans, maxLen int
	for i := 1; i <= limit; i++ {
		visited := make(map[int]bool)
		minMem := limit + 1
		j := i
		var len int
		for {
			t := sumDiv[j]
			minMem = tools.MinInt(minMem, t)
			if t == 0 || t > limit || visited[t] {
				break
			}
			if t == i {
				if maxLen == 0 || len > maxLen {
					ans = minMem
					maxLen = len
				}
				break
			}
			visited[t] = true
			j = t
			len++
		}
	}
	return ans
}

func main() {
	fmt.Println(solve(1000000))
}

// Find the smallest member of the longest amicable chain with no element
// exceeding one million.
// Note:
// Use the sieve of prime idea to compute all the sum of the proper divisors of
// all numbers. Then count the length of every possible amicable chain. And thus
// get the longest chain.
