package main

import "testing"

func TestP59(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{"./p059_cipher.txt", 107359},
	}
	for _, c := range cases {
		v := solve(c.in)
		if v != c.out {
			t.Errorf("P59: %v\tExpected: %v", v, c.out)
		}
	}
}
