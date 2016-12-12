package main

import "testing"

func Test89(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{"p089_roman.txt", 743},
	}
	for _, c := range cases {
		v := solve(c.in)
		if v != c.out {
			t.Errorf("p89: %v\tExpected: %v", v, c.out)
		}
	}
}
