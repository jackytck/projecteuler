package main

import "testing"

func TestP81(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{"./p081_matrix_small.txt", 2427},
		{"./p081_matrix.txt", 427337},
	}
	for _, c := range cases {
		v := solve(c.in)
		if v != c.out {
			t.Errorf("P81: %v\tExpected: %v", v, c.out)
		}
	}
}
