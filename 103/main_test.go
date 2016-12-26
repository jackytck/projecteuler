package main

import "testing"

func TestP103(t *testing.T) {
	ans := "20313839404245"
	v := solve()
	if v != ans {
		t.Errorf("p103: %v\tExpected: %v", v, ans)
	}
}
