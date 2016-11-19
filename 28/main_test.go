package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP28(t *testing.T) {
	cases := []tools.TestCase{
		{In: 5, Out: 101},
		{In: 1001, Out: 669171001},
	}
	tools.TestIntInt(t, cases, sumSpiralDiag, "P28")
}
