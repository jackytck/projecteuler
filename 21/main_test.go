package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP21(t *testing.T) {
	cases := []tools.TestCase{
		{In: 10000, Out: 31626},
	}
	tools.TestIntInt(t, cases, sumAmicable, "P21")
}
