package main

import (
	"testing"

	"github.com/jackytck/projecteuler/tools"
)

func TestP76(t *testing.T) {
	cases := []tools.TestCase{
		{In: 100, Out: 190569291},
	}
	tools.TestIntInt(t, cases, dp, "P76")
}
