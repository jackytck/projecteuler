package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP25(t *testing.T) {
	cases := []tools.TestCase{
		{In: 3, Out: 12},
		{In: 1000, Out: 4782},
	}
	tools.TestIntInt(t, cases, firstOver, "P25")
}
