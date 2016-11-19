package main

import (
	"fmt"
	"log"
	"strings"
)

func itowords(n int) string {
	if n > 1000 {
		log.Fatal(n, "is larger than 1000")
	}
	table := map[int]string{
		1:    "one",
		2:    "two",
		3:    "three",
		4:    "four",
		5:    "five",
		6:    "six",
		7:    "seven",
		8:    "eight",
		9:    "nine",
		10:   "ten",
		11:   "eleven",
		12:   "twelve",
		13:   "thirteen",
		14:   "fourteen",
		15:   "fifteen",
		16:   "sixteen",
		17:   "seventeen",
		18:   "eighteen",
		19:   "nineteen",
		20:   "twenty",
		30:   "thirty",
		40:   "forty",
		50:   "fifty",
		60:   "sixty",
		70:   "seventy",
		80:   "eighty",
		90:   "ninety",
		100:  "hundred",
		1000: "thousand",
	}

	if n == 1000 {
		return table[1] + " " + table[1000]
	}

	if v, ok := table[n]; ok && n != 100 {
		return v
	}

	if n >= 100 {
		d := n / 100
		if n%100 == 0 {
			return table[d] + " " + table[100]
		}
		return fmt.Sprintf("%s %s and %s", table[d], table[100], itowords(n%100))
	}

	if n > 20 {
		d := n / 10
		return fmt.Sprintf("%s-%s", table[d*10], itowords(n%10))
	}

	return ""
}

func count(s string) int {
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "-", "", -1)
	return len(s)
}

func solve() int {
	var sum int
	for i := 1; i <= 1000; i++ {
		sum += count(itowords(i))
	}
	return sum
}

func main() {
	fmt.Println(solve())
}

// Total number of letters used in writing out 1 to 1000 inclusive.
