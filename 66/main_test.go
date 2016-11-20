package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP66(t *testing.T) {
	cases := []tools.TestCase{
		{In: 1000, Out: 661},
	}
	tools.TestIntInt(t, cases, solve, "P66")
}
