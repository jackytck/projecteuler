package main

import "testing"

func TestP43(t *testing.T) {
	v := solve()
	out := 16695334890
	if v != out {
		t.Errorf("P43: %v\tExpected: %v", v, out)
	}
}
