package main

import "testing"

func TestP74(t *testing.T) {
	cases := []struct {
		in1, in2 int
		out      int
	}{
		{1000000, 60, 402},
	}
	for _, c := range cases {
		v := count(c.in1, c.in2)
		if v != c.out {
			t.Errorf("P74: %v\tExpected: %v", v, c.out)
		}
	}
}
