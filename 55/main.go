package main

import (
	"fmt"
	"math/big"

	"github.com/jackytck/projecteuler/tools"
)

func main() {
	var cnt int
	for i := 1; i < 10000; i++ {
		lychrel := big.NewInt(int64(i))
		for j := 0; j < 50; j++ {
			lychrel.Add(lychrel, tools.ReverseIntBig(lychrel))
			if tools.IsPalindromeIntBig(lychrel) {
				cnt++
				break
			}
		}
	}
	fmt.Println(9999 - cnt)
}

// How many Lychrel numbers are there below ten-thousand?
