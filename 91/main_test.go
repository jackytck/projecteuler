package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP92(t *testing.T) {
	cases := []tools.TestCase{
		{In: 2, Out: 14},
		{In: 50, Out: 14234},
	}
	tools.TestIntInt(t, cases, solve, "P92")
}
