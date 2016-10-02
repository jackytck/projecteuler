package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func read(path string) [][]int {
	lines, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	ss := strings.Split(string(lines), "\n")
	row := len(ss) - 1
	if row == 0 {
		log.Fatal()
	}
	grid := make([][]int, row)
	for i := 0; i < row; i++ {
		col := len(strings.Split(ss[i], " "))
		grid[i] = make([]int, col)
	}
	for i, v := range ss {
		for j, u := range strings.Split(v, " ") {
			if k, err := strconv.Atoi(u); err == nil {
				grid[i][j] = k
			}
		}
	}
	return grid
}

func min(a ...int) int {
	var ret int
	for i, v := range a {
		if i == 0 {
			ret = v
		}
		if v < ret {
			ret = v
		}
	}
	return ret
}

func max(a ...int) int {
	var ret int
	for i, v := range a {
		if i == 0 {
			ret = v
		}
		if v > ret {
			ret = v
		}
	}
	return ret
}

func dp(triangle [][]int) int {
	for i, u := range triangle {
		if i == 0 {
			continue
		}
		for j, v := range u {
			triangle[i][j] = max(triangle[i-1][max(0, j-1)], triangle[i-1][min(i-1, j)]) + v
		}
	}
	return max(triangle[len(triangle)-1]...)
}

func main() {
	small := read("./input_small.txt")
	large := read("./input.txt")
	fmt.Println(dp(small))
	fmt.Println(dp(large))
}

// Maximum path sum going from top to bottom of a triangle.
// Note: classic DP
