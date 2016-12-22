package main

import "testing"

func TestP99(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{"p099_base_exp.txt", 709},
	}
	for _, c := range cases {
		v := solve(c.in)
		if v != c.out {
			t.Errorf("P99: %v\tExpected: %v", v, c.out)
		}
	}
}
