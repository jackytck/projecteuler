package main

import "testing"

func TestP13(t *testing.T) {
	v := solve("./input.txt")
	out := 5537376230
	if v != out {
		t.Errorf("P13: %v\tExpected: %v", v, out)
	}
}
