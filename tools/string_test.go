package tools

import (
	"sort"
	"testing"
)

func TestSetString(t *testing.T) {
	cases := []struct {
		in  []string
		out []string
	}{
		{[]string{"a", "a", "b", "c", "a"}, []string{"a", "b", "c"}},
		{[]string{"@", "ab", "b", "cd", "b", "ab"}, []string{"b", "ab", "cd", "@"}},
	}
	for _, c := range cases {
		d := SetString(c.in)
		sort.Strings(d)
		sort.Strings(c.out)
		match := true
		for i, v := range d {
			if v != c.out[i] {
				match = false
				break
			}
		}
		if !match {
			t.Errorf("SetString: %v\tExpected: %v", d, c.out)
		}
	}
}
