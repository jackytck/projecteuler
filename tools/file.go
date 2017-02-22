package tools

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// ReadFile reads content of a file in a given path and returns it.
func ReadFile(path string) ([]string, error) {
	lines, err := ioutil.ReadFile(path)
	if err != nil {
		return []string{}, err
	}
	ss := strings.Split(string(lines), "\n")
	if ss[len(ss)-1] == "" {
		ss = ss[:len(ss)-1]
	}
	return ss, nil
}

// ReadWords reads content of a file that contains a line of words:
// "XX", "YYY", "ZZZ"
func ReadWords(path string) ([]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return []string{}, err
	}
	line := strings.Replace(string(data), "\"", "", -1)
	return strings.Split(string(line), ","), nil
}

// ReadMatrix reads an integer matrix stored in a file of a given path.
func ReadMatrix(path string) ([][]int, error) {
	var m [][]int
	lines, err := ioutil.ReadFile(path)
	if err != nil {
		return m, err
	}
	ss := strings.Split(string(lines), "\n")
	for i, row := range ss {
		if row == "" {
			continue
		}
		m = append(m, []int{})
		for _, v := range strings.Split(row, ",") {
			if v == "" {
				continue
			}
			d, _ := strconv.Atoi(v)
			m[i] = append(m[i], d)
		}
	}
	return m, nil
}
