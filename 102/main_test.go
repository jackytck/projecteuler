package main

import "testing"

func Test102(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{"p102_triangles.txt", 228},
	}
	for _, c := range cases {
		v := solve(c.in)
		if v != c.out {
			t.Errorf("p102: %v\tExpected: %v", v, c.out)
		}
	}
}
