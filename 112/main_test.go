package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP112(t *testing.T) {
	cases := []tools.TestCase{
		{In: 50, Out: 538},
		{In: 90, Out: 21780},
		{In: 99, Out: 1587000},
	}
	tools.TestIntInt(t, cases, solve, "P112")
}
