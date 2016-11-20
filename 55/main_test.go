package main

import "testing"

func TestP55(t *testing.T) {
	v := solve()
	out := 249
	if v != out {
		t.Errorf("P55: %v\tExpected: %v", v, out)
	}
}
