package main

import "testing"

func TestP5(t *testing.T) {
	cases := []struct {
		in  int
		out int
	}{
		{10, 2520},
		{20, 232792560},
	}
	for _, c := range cases {
		v := lcm(c.in)
		if v != c.out {
			t.Errorf("P5: %v\tExpected: %v", v, c.out)
		}
	}
}
