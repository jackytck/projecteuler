package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP6(t *testing.T) {
	cases := []tools.TestCase{
		{In: 10, Out: 2640},
		{In: 100, Out: 25164150},
	}
	tools.TestIntInt(t, cases, diff, "P6")
}
