package main

import "testing"

func Test100(t *testing.T) {
	cases := []struct {
		in  int64
		out int64
	}{
		{1000000000000, 756872327473},
	}
	for _, c := range cases {
		v := solve(c.in)
		if v != c.out {
			t.Errorf("p100: %v\tExpected: %v", v, c.out)
		}
	}
}
