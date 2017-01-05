package main

import "testing"

func TestP113(t *testing.T) {
	cases := []struct {
		in  int
		out int64
	}{
		{6, 12951},
		{10, 277032},
		{100, 51161058134250},
	}
	for _, c := range cases {
		v := solve(c.in).Int64()
		if v != c.out {
			t.Errorf("P113: %v\tExpected: %v", v, c.out)
		}
	}
}
