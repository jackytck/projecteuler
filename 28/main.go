package main

import "fmt"

func sumSpiralDiag(n int) int {
	return (4*n*n*n + 3*n*n + 8*n - 9) / 6
}

func main() {
	fmt.Println(sumSpiralDiag(5))
	fmt.Println(sumSpiralDiag(1001))
}

// Sum of numbers on the spiral diagonals.
// Note:
// The number of the top-right corner is always n^2.
// So the other 3 of the same level can be infered.
// Then summing all the levels is just simple maths.
