package main

import "testing"

func TestP27(t *testing.T) {
	v := solve()
	out := -59231
	if v != out {
		t.Errorf("P27: %v\tExpected: %v", v, out)
	}
}
