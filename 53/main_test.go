package main

import "testing"

func TestP53(t *testing.T) {
	v := solve()
	out := 4075
	if v != out {
		t.Errorf("P53: %v\tExpected: %v", v, out)
	}
}
