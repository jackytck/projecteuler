package tools

import (
	"io/ioutil"
	"os"
	"sort"
	"testing"
)

func TestReadFile(t *testing.T) {
	cases := []struct {
		in  string
		out []string
	}{
		{"first line\nsecond line\nthird line",
			[]string{"first line", "second line", "third line"}},
		{"first line\nsecond line\nthird line\n",
			[]string{"first line", "second line", "third line"}},
	}
	tmp := "tmp-input.txt"
	for _, c := range cases {
		ioutil.WriteFile(tmp, []byte(c.in), 0644)
		lines, e := ReadFile(tmp)
		if e != nil {
			t.Errorf("ReadFile: %v\tExpected: nil", e)
		}
		if len(lines) != len(c.out) {
			t.Errorf("ReadFile: len(%v)\tExpected: len(%v)", len(lines), len(c.out))
		}
		for i, v := range lines {
			if v != c.out[i] {
				t.Errorf("ReadFile: %v\tExpected: %v", v, c.out[i])
			}
		}
	}
	os.Remove(tmp)
	_, err := ReadFile("not_existed.txt")
	if err == nil {
		t.Error("ReadFile should return error if file does not exist.")
	}
}

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
		d, _ := ReadWords(tmp)
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
	_, err := ReadWords("not_existed.txt")
	if err == nil {
		t.Error("ReadWords should return error if file does not exist.")
	}
}
