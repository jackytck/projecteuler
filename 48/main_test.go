package main

import "testing"

func TestP48(t *testing.T) {
	cases := []struct {
		in1, in2 int
		out      int
	}{
		{10, 10, 405071317},
		{1000, 10, 9110846700},
	}
	for _, c := range cases {
		v := selfPower(c.in1, c.in2)
		if v != c.out {
			t.Errorf("P48: %v\tExpected: %v", v, c.out)
		}
	}
}
