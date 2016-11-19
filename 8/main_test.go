package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func wrap(a int) int {
	input := read("./input.txt")
	return largestProduct(input, a)
}

func TestP8(t *testing.T) {
	cases := []tools.TestCase{
		{In: 4, Out: 5832},
		{In: 13, Out: 23514624000},
	}
	tools.TestIntInt(t, cases, wrap, "P8")
}
