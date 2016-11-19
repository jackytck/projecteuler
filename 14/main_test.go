package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP14(t *testing.T) {
	cases := []tools.TestCase{
		{In: 1000000, Out: 837799},
	}
	tools.TestIntInt(t, cases, longest, "P14")
}
