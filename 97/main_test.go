package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP97(t *testing.T) {
	cases := []tools.TestCase{
		{In: 10, Out: 8739992577},
	}
	tools.TestIntInt(t, cases, solve, "P97")
}
