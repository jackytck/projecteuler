package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/jackytck/projecteuler/tools"
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

func dp(triangle [][]int) int {
	for i, u := range triangle {
		if i == 0 {
			continue
		}
		for j, v := range u {
			left := triangle[i-1][tools.MaxInt(0, j-1)]
			right := triangle[i-1][tools.MinInt(i-1, j)]
			triangle[i][j] = tools.MaxInt(left, right) + v
		}
	}
	return tools.MaxInt(triangle[len(triangle)-1]...)
}

func main() {
	small := read("./input_small.txt")
	large := read("./input.txt")
	fmt.Println(dp(small))
	fmt.Println(dp(large))
}

// Find the maximum total from top to bottom.
// Note:
// Classic DP.
// Same solution as problem 67.
