package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP39(t *testing.T) {
	cases := []tools.TestCase{
		{In: 1000, Out: 840},
	}
	tools.TestIntInt(t, cases, solve, "P39")
}
