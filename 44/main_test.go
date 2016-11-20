package main

import "testing"

func TestP44(t *testing.T) {
	v := fast()
	out := 5482660
	if v != out {
		t.Errorf("P44: %v\tExpected: %v", v, out)
	}
}
