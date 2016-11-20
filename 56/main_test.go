package main

import "testing"

func TestP56(t *testing.T) {
	v := solve()
	out := 972
	if v != out {
		t.Errorf("P56: %v\tExpected: %v", v, out)
	}
}
