package main

import "testing"

func TestP110(t *testing.T) {
	var ans int64 = 9350130049860600
	v := solve()
	if v != ans {
		t.Errorf("p110: %v\tExpected: %v", v, ans)
	}
}
