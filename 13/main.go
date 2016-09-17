package main

import (
	"fmt"
	"math/big"
)

func main() {
	sum := big.NewInt(0)
	num := big.NewInt(0)
	for {
		cnt, _ := fmt.Scan(num)
		if cnt == 0 {
			break
		}
		sum.Add(sum, num)
	}
	fmt.Println(sum.String()[:10])
}

// Find the first 10 digits of the sum of 100 50-digits numbers.
