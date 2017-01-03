package main

import "testing"

func Test111(t *testing.T) {
	cases := []struct {
		in  int
		out int64
	}{
		{4, 273700},
		{10, 612407567715},
	}
	for _, c := range cases {
		v := solve(c.in)
		if v != c.out {
			t.Errorf("p111: %v\tExpected: %v", v, c.out)
		}
	}
}
