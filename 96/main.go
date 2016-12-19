package main

import (
	"fmt"
	"strings"

	"github.com/jackytck/projecteuler/tools"
)

func solve(input string) int {
	var cnt int
	inputs := tools.ReadFile(input)
	for i, v := range inputs {
		if strings.HasPrefix(v, "Grid") {
			board := inputs[i+1 : i+10]
			su := Sudoku{}
			su.Init(board)
			su.Solve()
			cnt += su.Hash()
		} else {
			continue
		}
	}
	return cnt
}

func main() {
	fmt.Println(solve("./p096_sudoku.txt"))
}

// By solving all fifty puzzles, find the sum of the 3-digit numbers found in
// the top left corner of each solution grid.
// Note:
// Classic backtracking.
