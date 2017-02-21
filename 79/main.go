package main

import (
	"fmt"
	"log"

	"github.com/jackytck/projecteuler/tools"
)

func match(p, s string) bool {
	m := make(map[rune]bool)
	for _, r := range s {
		m[r] = true
	}
	var e string
	for _, r := range p {
		if m[r] {
			e += string(r)
		}
	}
	return e == s
}

func check(p string, log []string) bool {
	for _, s := range log {
		if !match(p, s) {
			return false
		}
	}
	return true
}

func solve(path string) string {
	// read
	logs, err := tools.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// make unique set
	exist := make(map[string]bool)
	var unique []string
	for _, v := range logs {
		if !exist[v] {
			unique = append(unique, v)
			exist[v] = true
		}
	}

	// assume no repeated digits and no digit 4 and 5
	for p := range tools.Perms([]int{0, 1, 2, 3, 6, 7, 8, 9}) {
		psw := tools.JoinIntsString(p...)
		if check(psw, unique) {
			return psw
		}
	}

	return ""
}

func main() {
	fmt.Println(solve("./p079_keylog.txt"))
}

// Given that the three characters are always asked for in order, analyse the
// file so as to determine the shortest possible secret passcode of unknown
// length.
// Note:
// Assume the passcode has no repeated character.
