package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP77(t *testing.T) {
	cases := []tools.TestCase{
		{In: 5000, Out: 71},
	}
	tools.TestIntInt(t, cases, dp, "P77")
}
