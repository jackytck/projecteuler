package main

import "testing"

func TestP62(t *testing.T) {
	v := solve()
	out := 127035954683
	if v != out {
		t.Errorf("P62: %v\tExpected: %v", v, out)
	}
}
