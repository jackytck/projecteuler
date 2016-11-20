package main

import "testing"

func TestP52(t *testing.T) {
	v := solve()
	out := 142857
	if v != out {
		t.Errorf("P52: %v\tExpected: %v", v, out)
	}
}
