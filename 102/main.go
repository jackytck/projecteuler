package main

import (
	"fmt"
	"log"

	"github.com/jackytck/projecteuler/tools"
)

func solve(input string) int {
	var cnt int
	lines, err := tools.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range lines {
		t := Triangle{}
		t.init(line)
		if t.isInside(0, 0) {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(solve("p102_triangles.txt"))
}

// Using the input, a 27K text file containing the co-ordinates of one thousand
// "random" triangles, find the number of triangles for which the interior
// contains the origin.
// Note:
// Barycentric coordinates is used to determine if a point lies inside a
// triangle:
// https://en.wikipedia.org/wiki/Barycentric_coordinate_system
