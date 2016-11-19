package main

import "testing"

func TestP18(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{"./input_small.txt", 23},
		{"./input.txt", 1074},
	}
	for _, c := range cases {
		v := dp(read(c.in))
		if v != c.out {
			t.Errorf("P18: %v\tExpected: %v", v, c.out)
		}
	}
}
