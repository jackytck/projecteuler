package main

import "testing"

func TestP63(t *testing.T) {
	v := solve()
	out := 49
	if v != out {
		t.Errorf("P63: %v\tExpected: %v", v, out)
	}
}
