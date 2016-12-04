package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func solve(path string) int {
	m := tools.ReadMatrix(path)
	for i, row := range m {
		for j := range row {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				m[i][j] += m[i][j-1]
			} else if j == 0 {
				m[i][j] += m[i-1][j]
			} else {
				m[i][j] += tools.MinInt(m[i-1][j], m[i][j-1])
			}
		}
	}
	s := len(m) - 1
	return m[s][s]
}

func main() {
	fmt.Println(solve("./p081_matrix_small.txt"))
	fmt.Println(solve("./p081_matrix.txt"))
}

// Find the minimal path sum of a matrix from the top left to the bottom right
// by only moving right and down.
// Note:
// Classic DP.
