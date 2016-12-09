package main

import "testing"

func Test86(t *testing.T) {
	cases := []struct {
		in1 int
		in2 int
		out int
	}{
		{2000, 200, 100},
		{1000000, 5000, 1818},
	}
	for _, c := range cases {
		v := solve(c.in1, c.in2)
		if v != c.out {
			t.Errorf("p86: %v\tExpected: %v", v, c.out)
		}
	}
}
