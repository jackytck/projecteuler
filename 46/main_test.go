package main

import "testing"

func TestP46(t *testing.T) {
	v := solve()
	out := 5777
	if v != out {
		t.Errorf("P46: %v\tExpected: %v", v, out)
	}
}
