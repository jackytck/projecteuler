package main

import "fmt"

func sum(a int) int {
	var ret int
	for i := 0; i < a; i++ {
		if i%3 == 0 || i%5 == 0 {
			ret += i
		}
	}
	return ret
}

func main() {
	fmt.Println(sum(10))
	fmt.Println(sum(1000))
}

// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
// The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.
