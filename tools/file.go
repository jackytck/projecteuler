package tools

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// ReadFile reads content of a file in a given path and returns it.
func ReadFile(path string) []string {
	lines, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	ss := strings.Split(string(lines), "\n")
	if ss[len(ss)-1] == "" {
		ss = ss[:len(ss)-1]
	}
	return ss
}

// ReadWords reads content of a file that contains a line of words:
// "XX", "YYY", "ZZZ"
func ReadWords(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	line := strings.Replace(string(data), "\"", "", -1)
	return strings.Split(string(line), ",")
}

// ReadMatrix reads an integer matrix stored in a file of a given path.
func ReadMatrix(path string) [][]int {
	var m [][]int
	lines, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
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
	return m
}
