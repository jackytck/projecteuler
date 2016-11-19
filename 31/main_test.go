package main

import "testing"

func TestP31(t *testing.T) {
	v := dp()
	out := 73682
	if v != out {
		t.Errorf("P31: %v\tExpected: %v", v, out)
	}
}
