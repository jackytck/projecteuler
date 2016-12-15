package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP92(t *testing.T) {
	cases := []tools.TestCase{
		{In: 10000000, Out: 8581146},
	}
	tools.TestIntInt(t, cases, solve, "P92")
}
