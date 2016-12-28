package main

import "testing"

func TestP105(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{"p105_sets.txt", 73702},
	}
	for _, c := range cases {
		v := solve(c.in)
		if v != c.out {
			t.Errorf("P105: %v\tExpected: %v", v, c.out)
		}
	}
}
