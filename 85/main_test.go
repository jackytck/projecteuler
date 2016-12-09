package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP85(t *testing.T) {
	cases := []tools.TestCase{
		{In: 2000000, Out: 2772},
	}
	tools.TestIntInt(t, cases, solve, "P85")
}
