package main

import "testing"

func TestP71(t *testing.T) {
	cases := []struct {
		in         int
		out1, out2 int
	}{
		{10, 2, 5},
		{1000000, 428570, 999997},
	}
	for _, c := range cases {
		v1, v2 := solve(c.in)
		if v1 != c.out1 || v2 != c.out2 {
			t.Errorf("p71: %v %v\tExpected: %v %v", v1, v2, c.out1, c.out2)
		}
	}
}
