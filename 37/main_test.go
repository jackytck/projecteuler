package main

import "testing"

func TestP37(t *testing.T) {
	v := solve()
	out := 748317
	if v != out {
		t.Errorf("P37: %v\tExpected: %v", v, out)
	}
}
