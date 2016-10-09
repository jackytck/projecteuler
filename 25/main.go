package main

import (
	"fmt"
	"math/big"
)

func firstOver(n int) int {
	a := big.NewInt(1)
	b := big.NewInt(1)
	i := 1
	for {
		if len(a.String()) == n {
			break
		}
		c := big.NewInt(0)
		c.Set(b)
		b.Add(a, b)
		a.Set(c)
		i++
	}
	return i
}

func main() {
	fmt.Println(firstOver(3))
	fmt.Println(firstOver(1000))
}

// The index of the first term in the Fibonacci sequence to contain 1000 digits.
