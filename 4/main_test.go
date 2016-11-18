package main

import "testing"

func TestP4(t *testing.T) {
	cases := []struct {
		in  int
		out int
	}{
		{2, 9009},
		{3, 906609},
	}
	for _, c := range cases {
		v, _, _ := largestPalindrome(c.in)
		if v != c.out {
			t.Errorf("P4: %v\tExpected: %v", v, c.out)
		}
	}
}
