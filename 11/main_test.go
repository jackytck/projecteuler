package main

import "testing"

func TestP11(t *testing.T) {
	v := solve()
	out := 70600674
	if v != out {
		t.Errorf("P11: %v\tExpected: %v", v, out)
	}
}
