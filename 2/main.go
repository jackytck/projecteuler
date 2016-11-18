package main

import "fmt"

func sum(limit int) int {
	i, j, sum := 0, 1, 0
	for i <= limit {
		i, j = j, i+j
		if i%2 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Println(sum(4000000))
}

// By considering the terms in the Fibonacci sequence whose values do not exceed
// four million, find the sum of the even-valued terms.
