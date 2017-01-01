package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP109(t *testing.T) {
	cases := []tools.TestCase{
		{In: 171, Out: 42336},
		{In: 100, Out: 38182},
	}
	tools.TestIntInt(t, cases, solve, "P109")
}
