package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var i int
	for v := range tools.Perms(a) {
		i++
		if i == 1000000 {
			for _, d := range v {
				fmt.Print(d)
			}
			fmt.Println()
			break
		}
	}
}

// The millionth lexicographic permutation of the digits: [0, 9].
