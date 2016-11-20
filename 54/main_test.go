package main

import "testing"

func TestP54(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{"./p054_poker.txt", 376},
	}
	for _, c := range cases {
		v := solve(c.in)
		if v != c.out {
			t.Errorf("P54: %v\tExpected: %v", v, c.out)
		}
	}
}
