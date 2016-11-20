package main

import "testing"

func TestP60(t *testing.T) {
	v := solve()
	out := 26033
	if v != out {
		t.Errorf("P60: %v\tExpected: %v", v, out)
	}
}
