package main

import "testing"

func TestP22(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{"./p022_names.txt", 871198282},
	}
	for _, c := range cases {
		v := solve(c.in)
		if v != c.out {
			t.Errorf("P22: %v\tExpected: %v", v, c.out)
		}
	}
}
