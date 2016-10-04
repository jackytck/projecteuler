package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func sumFactorialDigits(n int) int {
	var sum int
	z := big.NewInt(1)
	for i := 2; i <= n; i++ {
		z.Mul(z, big.NewInt(int64(i)))
	}
	for _, v := range z.String() {
		v, _ := strconv.Atoi(string(v))
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sumFactorialDigits(10))
	fmt.Println(sumFactorialDigits(100))
}

// Sum of the digits in the number n!.
