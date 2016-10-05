package tools

import "math"

// Sum returns the sum of its argument.
func Sum(a ...int) int {
	var s int
	for _, v := range a {
		s += v
	}
	return s
}

// SqrtInt computes the square root of n and casts to an int.
func SqrtInt(n int) int {
	return int(math.Sqrt(float64(n)))
}

// Divisors returns all the proper divisors of n.
func Divisors(n int) []int {
	d := []int{1}
	for i := 2; i <= SqrtInt(n); i++ {
		if n%i == 0 {
			d = append(d, i)
			if other := n / i; other != i {
				d = append(d, other)
			}
		}
	}
	return d
}

// SumDivisors sums all the proper divisors of n.
func SumDivisors(n int) int {
	return Sum(Divisors(n)...)
}

// Perms return a channel of each permutation of a slice.
func Perms(slice []int) chan []int {
	c := make(chan []int)
	go func() {
		if len(slice) == 1 {
			c <- slice
		} else {
			for i, v := range slice {
				b := append([]int(nil), slice[:i]...)
				b = append(b, slice[i+1:]...)
				for p := range Perms(b) {
					c <- append([]int{v}, p...)
				}
			}
		}
		close(c)
	}()
	return c
}

// Divmod gives the quotient and remainder of integer division.
func Divmod(a, b int) (int, int) {
	return a / b, a % b
}
