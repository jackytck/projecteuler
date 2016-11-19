package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP20(t *testing.T) {
	cases := []tools.TestCase{
		{In: 10, Out: 27},
		{In: 100, Out: 648},
	}
	tools.TestIntInt(t, cases, sumFactorialDigits, "P20")
}
