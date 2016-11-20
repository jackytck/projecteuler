package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP57(t *testing.T) {
	cases := []tools.TestCase{
		{In: 1000, Out: 153},
	}
	tools.TestIntInt(t, cases, count, "P57")
}
