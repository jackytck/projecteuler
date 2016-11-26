package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP73(t *testing.T) {
	cases := []tools.TestCase{
		{In: 8, Out: 3},
		{In: 12000, Out: 7295372},
	}
	tools.TestIntInt(t, cases, count, "P73")
}
