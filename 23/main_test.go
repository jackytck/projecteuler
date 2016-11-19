package main

import "testing"

func TestP23(t *testing.T) {
	v := solve()
	out := 4179871
	if v != out {
		t.Errorf("P23: %v\tExpected: %v", v, out)
	}
}
