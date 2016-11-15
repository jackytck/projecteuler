package main

import (
	"fmt"
	"math/big"
	"sort"

	"github.com/jackytck/projecteuler/tools"
)

// Node stores the number of times that a key could be hashed from the
// (smallest base)^3.
type Node struct {
	count int
	base  int
}

func hash(x int) int64 {
	t := int64(x)
	b := big.NewInt(t * t * t)
	d := tools.DigitsBig(b)
	sort.Sort(sort.Reverse(sort.IntSlice(d)))
	return tools.JoinIntsBig(d).Int64()
}

func main() {
	cubeMap := make(map[int64]Node)
	i := 346
	for {
		k := hash(i)
		n, ok := cubeMap[k]
		if ok {
			n.count++
			if i < n.base {
				n.base = i
			}
			cubeMap[k] = n
			if n.count == 5 {
				b := n.base
				fmt.Printf("%v = %v^3\n", b*b*b, b)
				break
			}
		} else {
			cubeMap[k] = Node{1, i}
		}
		i++
	}
}

// Find the smallest cube for which exactly five permutations of its digits are cube.
// Note:
// Assume the first answer we see has exactly five permutations, but not at least
// five.
