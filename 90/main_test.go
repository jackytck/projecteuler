package main

import "testing"

func TestP90(t *testing.T) {
	ans := 1217
	v := solve()
	if v != ans {
		t.Errorf("p90: %v\tExpected: %v", v, ans)
	}
}
