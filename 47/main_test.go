package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP47(t *testing.T) {
	cases := []tools.TestCase{
		{In: 2, Out: 14},
		{In: 3, Out: 644},
		{In: 4, Out: 134043},
	}
	tools.TestIntInt(t, cases, consecutive, "P47")
}
