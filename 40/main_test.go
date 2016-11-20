package main

import "testing"

func TestP40(t *testing.T) {
	v := solve()
	out := 210
	if v != out {
		t.Errorf("P40: %v\tExpected: %v", v, out)
	}
}
