package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP50(t *testing.T) {
	cases := []tools.TestCase{
		{In: 100, Out: 41},
		{In: 1000, Out: 953},
		{In: 1000000, Out: 997651},
	}
	tools.TestIntInt(t, cases, consecutivePrimeSum, "P50")
}
