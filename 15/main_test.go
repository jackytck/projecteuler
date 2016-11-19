package main

import "testing"

func TestP15(t *testing.T) {
	cases := []struct {
		in1, in2 int64
		out      int64
	}{
		{4, 2, 6},
		{40, 20, 137846528820},
	}
	for _, c := range cases {
		v := nCr(c.in1, c.in2)
		if v.Int64() != c.out {
			t.Errorf("P15: %v\tExpected: %v", v, c.out)
		}
	}
}
