package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP95(t *testing.T) {
	cases := []tools.TestCase{
		{In: 1000000, Out: 14316},
	}
	tools.TestIntInt(t, cases, solve, "P95")
}
