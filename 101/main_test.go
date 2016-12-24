package main

import "testing"

func TestP101(t *testing.T) {
	ans := 37076114526
	v := solve()
	if v != ans {
		t.Errorf("p101: %v\tExpected: %v", v, ans)
	}
}
