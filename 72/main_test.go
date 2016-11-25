package main

import "testing"

func Test72(t *testing.T) {
	cases := []struct {
		in  int
		out int64
	}{
		{8, 21},
		{1000000, 303963552391},
	}
	for _, c := range cases {
		v := solve(c.in)
		if v != c.out {
			t.Errorf("p72: %v\tExpected: %v", v, c.out)
		}
	}
}
