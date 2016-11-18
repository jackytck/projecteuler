package main

import "testing"

func TestP1(t *testing.T) {
	cases := []struct {
		in  int
		out int
	}{
		{10, 23},
		{1000, 233168},
	}
	for _, c := range cases {
		v := sum(c.in)
		if v != c.out {
			t.Errorf("P1: %v\tExpected: %v", v, c.out)
		}
	}
}
