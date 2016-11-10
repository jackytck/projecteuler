package main

import (
	"fmt"
	"math/big"

	"github.com/jackytck/projecteuler/tools"
)

func main() {
	var cnt int
	d := big.NewInt(3)
	n := big.NewInt(2)
	for i := 1; i < 1000; i++ {
		p := big.NewInt(0)
		p.Set(n)
		n.Add(d, n)
		d.Add(n, p)
		if len(tools.DigitsBig(d)) > len(tools.DigitsBig(n)) {
			cnt++
		}
	}
	fmt.Println(cnt)
}

// In the first one-thousand expansions of square root of 2,
// how many fractions contain a numerator with more digits than denominator?
