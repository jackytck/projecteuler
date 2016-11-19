package main

import "testing"

func TestP33(t *testing.T) {
	_, v := solve()
	out := 100
	if v != out {
		t.Errorf("P33: %v\tExpected: %v", v, out)
	}
}
