package main

import "testing"

func TestP88(t *testing.T) {
	ans := 7587457
	v := solve()
	if v != ans {
		t.Errorf("p88: %v\tExpected: %v", v, ans)
	}
}
