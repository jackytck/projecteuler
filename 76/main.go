package main

import "fmt"

func dp(goal int) int {
	table := make(map[int]int)
	table[0] = 1
	for c := 1; c < goal; c++ {
		for i := c; i <= goal; i++ {
			table[i] += table[i-c]
		}
	}
	return table[goal]
}

func main() {
	fmt.Println(dp(100))
}

// How many different ways can one hundred be written as a sum of at least two
// positive integers?
// Note:
// This is the same DP as problem 31. The only difference is: now we are
// considering coin 1, 2, 3, 4, ..., 98, 99, and want to find the total number
// of ways that could pay 100 dollars.
