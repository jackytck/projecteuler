package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP16(t *testing.T) {
	cases := []tools.TestCase{
		{In: 15, Out: 26},
		{In: 1000, Out: 1366},
	}
	tools.TestIntInt(t, cases, powerDigitSum, "P16")
}
