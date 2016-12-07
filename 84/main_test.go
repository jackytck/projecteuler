package main

import "testing"

func Test84(t *testing.T) {
	cases := []struct {
		in1, in2 int
		out      string
	}{
		{10000000, 6, "102400"},
		{10000000, 4, "101524"},
	}
	for _, c := range cases {
		v := monteCarlo(c.in1, c.in2)
		if v != c.out {
			t.Errorf("p84: %v\tExpected: %v", v, c.out)
		}
	}
}
