package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP7(t *testing.T) {
	cases := []tools.TestCase{
		{In: 6, Out: 13},
		{In: 10001, Out: 104743},
	}
	tools.TestIntInt(t, cases, nthPrime, "P7")
}
