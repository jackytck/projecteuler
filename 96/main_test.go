package main

import "testing"

func Test96(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{"./p096_sudoku.txt", 24702},
	}
	for _, c := range cases {
		v := solve(c.in)
		if v != c.out {
			t.Errorf("p96: %v\tExpected: %v", v, c.out)
		}
	}
}
