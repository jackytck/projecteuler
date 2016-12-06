package main

import "testing"

func TestP83(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{"./p083_matrix_small.txt", 2297},
		{"./p083_matrix.txt", 425185},
	}
	for _, c := range cases {
		v := solve(c.in)
		if v != c.out {
			t.Errorf("P83: %v\tExpected: %v", v, c.out)
		}
	}
}
