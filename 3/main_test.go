package main

import "testing"

func TestP3(t *testing.T) {
	cases := []struct {
		in  uint64
		out uint64
	}{
		{13195, 29},
		{600851475143, 6857},
	}
	for _, c := range cases {
		v := largestPrime(c.in)
		if v != c.out {
			t.Errorf("P3: %v\tExpected: %v", v, c.out)
		}
	}
}
