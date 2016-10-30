package tools

import (
	"io/ioutil"
	"log"
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
