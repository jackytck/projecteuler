package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP41(t *testing.T) {
	cases := []tools.TestCase{
		{In: 7654321, Out: 7652413},
	}
	tools.TestIntInt(t, cases, pandigitalPrime, "P41")
}
