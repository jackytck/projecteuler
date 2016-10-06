package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func nameScore(n string) int {
	var sum int
	for _, v := range n {
		sum += int(v - 'A' + 1)
	}
	return sum
}

func read(path string) []string {
	lines, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	ss := strings.Split(string(lines), ",")
	for i := 0; i < len(ss); i++ {
		ss[i] = strings.Replace(ss[i], "\"", "", -1)
	}
	return ss
}

func main() {
	var score int
	names := read("./p022_names.txt")
	sort.Strings(names)
	for i, v := range names {
		score += (i + 1) * nameScore(v)
	}
	fmt.Println(score)
}

// Total name scores of names in a file.
