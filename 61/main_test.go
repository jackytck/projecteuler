package main

import "testing"

func TestP61(t *testing.T) {
	v := solve()
	out := 28684
	if v != out {
		t.Errorf("P61: %v\tExpected: %v", v, out)
	}
}
