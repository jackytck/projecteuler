package main

import "testing"

func TestP98(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{"p098_words.txt", 18769},
	}
	for _, c := range cases {
		v := solve(c.in)
		if v != c.out {
			t.Errorf("P98: %v\tExpected: %v", v, c.out)
		}
	}
}
