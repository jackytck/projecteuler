package main

import "testing"

func TestP30(t *testing.T) {
	v := solve()
	out := 443839
	if v != out {
		t.Errorf("P30: %v\tExpected: %v", v, out)
	}
}
