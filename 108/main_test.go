package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP108(t *testing.T) {
	cases := []tools.TestCase{
		{In: 2, Out: 4},
		{In: 1000, Out: 180180},
	}
	tools.TestIntInt(t, cases, solve, "P108")
}
