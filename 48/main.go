package main

import (
	"fmt"
	"math/big"

	"github.com/jackytck/projecteuler/tools"
)

func selfPower(n, last int) int {
	var sum int
	m := tools.Exp(10, last)
	for i := 1; i <= n; i++ {
		z := big.NewInt(int64(i))
		z.Exp(z, z, m)
		sum += int(z.Int64())
	}
	return sum % int(m.Int64())
}

func main() {
	fmt.Printf("%010d\n", selfPower(10, 10))
	fmt.Printf("%010d\n", selfPower(1000, 10))
}

// Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.
