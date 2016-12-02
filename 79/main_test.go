package main

import "testing"

func Test79(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{"./p079_keylog.txt", "73162890"},
	}
	for _, c := range cases {
		v := solve(c.in)
		if v != c.out {
			t.Errorf("p79: %v\tExpected: %v", v, c.out)
		}
	}
}
