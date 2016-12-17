package main

import "testing"

func Test94(t *testing.T) {
	cases := []struct {
		in  int64
		out int64
	}{
		{1000000000, 518408346},
	}
	for _, c := range cases {
		v := solve(c.in)
		if v != c.out {
			t.Errorf("p94: %v\tExpected: %v", v, c.out)
		}
	}
}
