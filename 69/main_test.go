package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP69(t *testing.T) {
	cases := []tools.TestCase{
		{In: 1000000, Out: 510510},
	}
	tools.TestIntInt(t, cases, solve, "P69")
}
