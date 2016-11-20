package main

import "testing"

func TestP38(t *testing.T) {
	v := solve()
	out := 932718654
	if v != out {
		t.Errorf("P38: %v\tExpected: %v", v, out)
	}
}
