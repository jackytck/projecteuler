package main

import "testing"

func TestP32(t *testing.T) {
	v := solve()
	out := 45228
	if v != out {
		t.Errorf("P32: %v\tExpected: %v", v, out)
	}
}
