package main

import "fmt"

func sumOfSquare(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func sumOfFirst(n int) int {
	return n * (n + 1) / 2
}

func squareOfSum(n int) int {
	s := sumOfFirst(n)
	return s * s
}

func diff(n int) int {
	return squareOfSum(n) - sumOfSquare(n)
}

func main() {
	fmt.Println(diff(10))
	fmt.Println(diff(100))
}

// Difference between the sum of the squares and the square of the sum.
