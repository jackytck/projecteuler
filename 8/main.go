package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func read(path string) string {
	digit, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	s := string(digit)
	s = strings.Replace(s, "\n", "", -1)
	return s
}

func multiply(input string) int {
	p := 1
	for _, v := range input {
		d, _ := strconv.Atoi(string(v))
		p *= d
	}
	return p
}

func largestProduct(input string, length int) int {
	lp := 1
	for i := 0; i <= len(input)-length; i++ {
		s := input[i : i+length]
		if p := multiply(s); p > lp {
			lp = p
		}
	}
	return lp
}

func main() {
	input := read("./input.txt")
	fmt.Println(largestProduct(input, 4))
	fmt.Println(largestProduct(input, 13))
}

// N adjacent digits in a given number that have the greatest product.
