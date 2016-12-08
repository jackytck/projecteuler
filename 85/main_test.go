package main

import "testing"

func Test85(t *testing.T) {
	cases := []struct {
		in  int
		out int
	}{
		{2000000, 2772},
	}
	for _, c := range cases {
		v := solve(c.in)
		if v != c.out {
			t.Errorf("p85: %v\tExpected: %v", v, c.out)
		}
	}
}
