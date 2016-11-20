package main

import "testing"

func TestP51(t *testing.T) {
	v := solve()
	out := 121313
	if v != out {
		t.Errorf("P51: %v\tExpected: %v", v, out)
	}
}
