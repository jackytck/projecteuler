package main

import (
	"fmt"
	"math"
)

func main() {
	var cnt int
	n := 1
	var t int
	for t < 10 {
		t = int(math.Ceil(math.Pow(10, float64(n-1)/float64(n))))
		cnt += 10 - t
		n++
	}
	fmt.Println(cnt)
}

// How many n-digit positive integers exist which are also an nth power?
// Note:
// Equivalent of asking the number of integer solution of x for which:
// 10^(n-1) <= x^n < 10^n
// So we could get the range for x and n.
