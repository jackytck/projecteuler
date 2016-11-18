package main

import "testing"

func TestP2(t *testing.T) {
	cases := []struct {
		in  int
		out int
	}{
		{4000000, 4613732},
	}
	for _, c := range cases {
		v := sum(c.in)
		if v != c.out {
			t.Errorf("P2: %v\tExpected: %v", v, c.out)
		}
	}
}
