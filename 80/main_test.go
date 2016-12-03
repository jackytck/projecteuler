package main

import "testing"

func TestP80(t *testing.T) {
	cases := []struct {
		in  int
		out int64
	}{
		{100, 40886},
	}
	for _, c := range cases {
		v := solve(c.in)
		if v != c.out {
			t.Errorf("P80: %v\tExpected: %v", v, c.out)
		}
	}
}
