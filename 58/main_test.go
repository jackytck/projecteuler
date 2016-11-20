package main

import "testing"

func TestP58(t *testing.T) {
	v := solve()
	out := 26241
	if v != out {
		t.Errorf("P58: %v\tExpected: %v", v, out)
	}
}
