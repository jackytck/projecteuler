package tools

import (
	"io/ioutil"
	"os"
	"sort"
	"testing"
)

func TestReadWords(t *testing.T) {
	cases := []struct {
		in  string
		out []string
	}{
		{`"A","ABILITY","ABLE","ABOUT","ABOVE","ABSENCE","ABSOLUTELY","ACADEMIC"`,
			[]string{"A", "ABILITY", "ABLE", "ABOUT", "ABOVE", "ABSENCE", "ABSOLUTELY", "ACADEMIC"}},
	}
	tmp := "tmp-input.txt"
	for _, c := range cases {
		ioutil.WriteFile(tmp, []byte(c.in), 0644)
		d := ReadWords(tmp)
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
			t.Errorf("ReadWords: %v\tExpected: %v", d, c.out)
		}
	}
	os.Remove(tmp)
}
