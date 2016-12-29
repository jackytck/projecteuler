package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP106(t *testing.T) {
	cases := []tools.TestCase{
		{In: 4, Out: 1},
		{In: 7, Out: 70},
		{In: 12, Out: 21384},
	}
	tools.TestIntInt(t, cases, solve, "P106")
}
