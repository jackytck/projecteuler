package main

import "testing"

func TestP45(t *testing.T) {
	v := solve()
	out := 1533776805
	if v != out {
		t.Errorf("P45: %v\tExpected: %v", v, out)
	}
}
