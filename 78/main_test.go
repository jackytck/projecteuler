package main

import "testing"

func TestP78(t *testing.T) {
	v := solve()
	out := 55374
	if v != out {
		t.Errorf("P78: %v\tExpected: %v", v, out)
	}
}
