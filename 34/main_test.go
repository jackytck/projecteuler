package main

import "testing"

func TestP34(t *testing.T) {
	v := solve()
	out := 40730
	if v != out {
		t.Errorf("P34: %v\tExpected: %v", v, out)
	}
}
