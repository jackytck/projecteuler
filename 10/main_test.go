package main

import "testing"

func TestP10(t *testing.T) {
	cases := []struct {
		in  int
		out uint64
	}{
		{10, 17},
		{2000000, 142913828922},
	}
	for _, c := range cases {
		v := sumPrimes(c.in)
		if v != c.out {
			t.Errorf("P10: %v\tExpected: %v", v, c.out)
		}
	}
}
