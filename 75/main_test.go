package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP75(t *testing.T) {
	cases := []tools.TestCase{
		{In: 1500000, Out: 161667},
	}
	tools.TestIntInt(t, cases, count, "P75")
}
