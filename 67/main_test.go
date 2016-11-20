package main

import "testing"

func TestP67(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{"./p067_triangle.txt", 7273},
	}
	for _, c := range cases {
		v := dp(read(c.in))
		if v != c.out {
			t.Errorf("P67: %v\tExpected: %v", v, c.out)
		}
	}
}
