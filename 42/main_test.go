package main

import "testing"

func TestP42(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{"./p042_words.txt", 162},
	}
	for _, c := range cases {
		v := solve(c.in)
		if v != c.out {
			t.Errorf("P42: %v\tExpected: %v", v, c.out)
		}
	}
}
