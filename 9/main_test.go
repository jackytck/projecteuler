package main

import "testing"

func TestP9(t *testing.T) {
	v := solve()
	out := 31875000
	if v != out {
		t.Errorf("%s: %v\tExpected: %v", "P9", v, out)
	}
}
