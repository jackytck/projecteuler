package main

import "fmt"

func collatz(n int) int {
	var length int
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		length++
	}
	return length + 1
}

func longest(n int) int {
	var start, max int
	for i := 1; i < n; i++ {
		if c := collatz(i); c > max {
			start = i
			max = c
		}
	}
	return start
}

func main() {
	fmt.Println(longest(1000000))
}

// The starting number (under a million) of a Collatz sequence that produces the longest chain.
