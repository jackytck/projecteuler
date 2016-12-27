package main

import "testing"

func TestP104(t *testing.T) {
	ans := 329468
	v := solve()
	if v != ans {
		t.Errorf("p104: %v\tExpected: %v", v, ans)
	}
}
