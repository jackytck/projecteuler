package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP87(t *testing.T) {
	cases := []tools.TestCase{
		{In: 50, Out: 4},
		{In: 50000000, Out: 1097343},
	}
	tools.TestIntInt(t, cases, solve, "P87")
}
