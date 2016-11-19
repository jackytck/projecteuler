package main

import "testing"

func TestP29(t *testing.T) {
	cases := []struct {
		in1, in2 int
		out      int
	}{
		{2, 5, 15},
		{2, 100, 9183},
	}
	for _, c := range cases {
		v := distinct(c.in1, c.in2)
		if v != c.out {
			t.Errorf("P29: %v\tExpected: %v", v, c.out)
		}
	}
}
