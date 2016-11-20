package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP65(t *testing.T) {
	cases := []tools.TestCase{
		{In: 100, Out: 272},
	}
	tools.TestIntInt(t, cases, solve, "P65")
}
