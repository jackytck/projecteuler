package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP64(t *testing.T) {
	cases := []tools.TestCase{
		{In: 10000, Out: 1322},
	}
	tools.TestIntInt(t, cases, solve, "P64")
}
