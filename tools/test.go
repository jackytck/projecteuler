package tools

import "testing"

// TestCase represents a simple int in int out test case.
type TestCase struct {
	In  int
	Out int
}

// TestIntInt tests a simple func(int) int function.
func TestIntInt(t *testing.T, cases []TestCase, f func(int) int, name string) {
	for _, c := range cases {
		v := f(c.In)
		if v != c.Out {
			t.Errorf("%s: %v\tExpected: %v", name, v, c.Out)
		}
	}
}
