package main

import (
	"fmt"
	"math"

	"github.com/jackytck/projecteuler/tools"
)

func champernowne(n int) int {
	var s int
	var i int
	var p int
	var d int
	for s < n {
		d = int(math.Pow10(i))
		p = 9 * d * (i + 1)
		s += p
		i++
	}
	q, r := tools.Divmod(n-s+p, i)
	if r == 0 {
		return (d + q - 1) % 10
	}
	return tools.DigitsIth(d+q, r-1)
}

func main() {
	p := 1
	for i := 0; i <= 6; i++ {
		p *= champernowne(int(math.Pow10(i)))
	}
	fmt.Println(p)
}

// Find the product of digits of d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000,
// where di is the i-th digit of a champernowne's constant.
