package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func exp2(n uint) *big.Int {
	b := big.NewInt(1)
	b.Lsh(b, n)
	return b
}

func powerDigitSum(n int) int {
	sum := 0
	p := exp2(uint(n))
	for _, v := range p.Text(10) {
		i, _ := strconv.Atoi(string(v))
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(powerDigitSum(15))
	fmt.Println(powerDigitSum(1000))
}

// Sum of digits of 2^n.
