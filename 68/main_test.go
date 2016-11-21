package main

import "testing"

func TestP68(t *testing.T) {
	v := solve()
	out := "6531031914842725"
	if v != out {
		t.Errorf("P68: %v\tExpected: %v", v, out)
	}
}
