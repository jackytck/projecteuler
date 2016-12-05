package main

import "testing"

func TestP82(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{"./p082_matrix_small.txt", 994},
		{"./p082_matrix.txt", 260324},
	}
	for _, c := range cases {
		v := solve(c.in)
		if v != c.out {
			t.Errorf("P82: %v\tExpected: %v", v, c.out)
		}
	}
}
