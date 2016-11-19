package tools

import "testing"

func TestGCD(t *testing.T) {
	cases := []struct {
		in1, in2 int
		out      int
	}{
		{0, 0, -1},
		{8, 12, 4},
		{54, 24, 6},
		{18, 48, 6},
		{42, 56, 14},
		{-35, 280, 35},
		{252, 105, 21},
	}
	for _, c := range cases {
		v := GCD(c.in1, c.in2)
		if v != c.out {
			t.Errorf("GCP: %v\tExpected: %v", v, c.out)
		}
	}
}
