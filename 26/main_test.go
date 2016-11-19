package main

import "testing"

func TestP26(t *testing.T) {
	v := solve()
	out := 983
	if v != out {
		t.Errorf("P26: %v\tExpected: %v", v, out)
	}
}
