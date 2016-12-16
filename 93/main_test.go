package main

import "testing"

func TestP93(t *testing.T) {
	ans := 1258
	v := solve()
	if v != ans {
		t.Errorf("p93: %v\tExpected: %v", v, ans)
	}
}
