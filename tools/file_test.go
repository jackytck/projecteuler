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

func TestReadMatrix(t *testing.T) {
	cases := []struct {
		in  string
		out [][]int
	}{
		{"131,673,234,103,18\n201,96,342,965,150\n" +
			"630,803,746,422,111\n537,699,497,121,956\n" +
			"805,732,524,37,331",
			[][]int{{131, 673, 234, 103, 18}, {201, 96, 342, 965, 150},
				{630, 803, 746, 422, 111}, {537, 699, 497, 121, 956},
				{805, 732, 524, 37, 331}}},
		{"131,673,234,103,18\n201,96,342,965,150\n",
			[][]int{{131, 673, 234, 103, 18}, {201, 96, 342, 965, 150}}},
		{"131,673,234,103,18\n201,96,342,965,150,\n",
			[][]int{{131, 673, 234, 103, 18}, {201, 96, 342, 965, 150}}},
	}
	tmp := "tmp-input.txt"
	for _, c := range cases {
		ioutil.WriteFile(tmp, []byte(c.in), 0644)
		mat, e := ReadMatrix(tmp)
		if e != nil {
			t.Errorf("ReadMatrix: %v\tExpected: nil", e)
		}
		if len(mat) != len(c.out) {
			t.Errorf("ReadMatrix: len(%v)\tExpected: len(%v)", len(mat), len(c.out))
		}
		for i, v := range mat {
			if len(v) != len(c.out[i]) {
				t.Errorf("ReadMatrix: len(%v)\tExpected: len(%v)", len(v), len(c.out[i]))
			}
			for j, u := range v {
				if u != c.out[i][j] {
					t.Errorf("ReadMatrix: %v\tExpected: %v", u, c.out[i][j])
				}
			}
		}
	}
	os.Remove(tmp)
	_, err := ReadMatrix("not_existed.txt")
	if err == nil {
		t.Error("ReadMatrix should return error if file does not exist.")
	}
}
