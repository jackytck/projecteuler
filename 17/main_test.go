package main

import "testing"

func TestP17(t *testing.T) {
	v := solve()
	out := 21124
	if v != out {
		t.Errorf("P17: %v\tExpected: %v", v, out)
	}
}
