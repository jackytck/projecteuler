package main

import "testing"

func TestP49(t *testing.T) {
	v := solve()
	out := "296962999629"
	if v != out {
		t.Errorf("P49: %v\tExpected: %v", v, out)
	}
}
