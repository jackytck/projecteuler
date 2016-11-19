package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP12(t *testing.T) {
	cases := []tools.TestCase{
		{In: 5, Out: 28},
		{In: 500, Out: 76576500},
	}
	tools.TestIntInt(t, cases, search, "P12")
}
