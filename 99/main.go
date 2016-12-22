package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jackytck/projecteuler/tools"
)

func solve(input string) int {
	var maxI int
	var maxV float64
	for i, v := range tools.ReadFile(input) {
		p := strings.Split(v, ",")
		a, _ := strconv.Atoi(p[0])
		b, _ := strconv.Atoi(p[1])
		if f := float64(b) * math.Log(float64(a)); maxI == 0 || f > maxV {
			maxI, maxV = i+1, f
		}
	}
	return maxI
}

func main() {
	fmt.Println(solve("p099_base_exp.txt"))
}

// Using base_exp.txt, a 22K text file containing one thousand lines with a
// base/exponent pair on each line, determine which line number has the greatest
// numerical value.
// Note:
// Just compare a * ln(b), assume no floating point precision problem.
