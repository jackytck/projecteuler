package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP36(t *testing.T) {
	cases := []tools.TestCase{
		{In: 1000000, Out: 872187},
	}
	tools.TestIntInt(t, cases, solve, "P36")
}
