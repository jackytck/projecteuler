package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP35(t *testing.T) {
	cases := []tools.TestCase{
		{In: 100, Out: 13},
		{In: 1000000, Out: 55},
	}
	tools.TestIntInt(t, cases, circularPrimes, "P35")
}
