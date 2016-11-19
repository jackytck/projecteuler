package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP24(t *testing.T) {
	cases := []tools.TestCase{
		{In: 1000000, Out: 2783915460},
	}
	tools.TestIntInt(t, cases, solve, "P24")
}
