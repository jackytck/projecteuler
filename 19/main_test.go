package main

import "testing"

func TestP19(t *testing.T) {
	v := countWeekday(1901, 2000, 0)
	out := 171
	if v != out {
		t.Errorf("P19: %v\tExpected: %v", v, out)
	}
}
