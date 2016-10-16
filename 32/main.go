package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func main() {
	var sum int
	digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	pan := make(map[int]bool)
	for p := range tools.Perms(digits) {
		for i := 1; i < 7; i++ {
			for j := i + 1; j < 8; j++ {
				a := tools.JoinInts(p[:i])
				b := tools.JoinInts(p[i:j])
				c := tools.JoinInts(p[j:])

				if a*b == c {
					_, ok := pan[c]
					if ok {
						continue
					}
					pan[c] = true
					sum += c
				}
			}
		}
	}
	fmt.Println(sum)
}

// Sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.
