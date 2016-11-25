package main

import "testing"

func TestP70(t *testing.T) {
	ans := 8319823
	v := solve()
	if v != ans {
		t.Errorf("p70: %v\tExpected: %v", v, ans)
	}
}
