package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

func read(input string) string {
	lines, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	return string(lines)
}

func solve(input string) int {
	roman := read(input)
	re := regexp.MustCompile("DCCCC|LXXXX|VIIII|CCCC|XXXX|IIII")
	minimal := re.ReplaceAllLiteralString(roman, "--")
	return len(roman) - len(minimal)
}

func main() {
	fmt.Println(solve("p089_roman.txt"))
}

// Input file contains one thousand numbers written in valid, but not
// necessarily minimal, Roman numerals. Find the number of characters saved by
// writing each of these in their minimal form.
// Note:
// Just replace each of the 6 long forms into a 2-characters subtractive form.
