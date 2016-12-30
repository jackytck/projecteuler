package main

import "testing"

func TestP107(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{"p107_network.txt", 259679},
	}
	for _, c := range cases {
		v := solve(c.in)
		if v != c.out {
			t.Errorf("P107: %v\tExpected: %v", v, c.out)
		}
	}
}
