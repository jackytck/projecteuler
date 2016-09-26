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
	col := len(strings.Split(ss[0], " "))
	grid := make([][]int, row)
	for i := 0; i < row; i++ {
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

func transpose(grid [][]int) [][]int {
	row := len(grid)
	if row == 0 {
		return nil
	}
	col := len(grid[0])
	trans := make([][]int, col)
	for i := 0; i < col; i++ {
		trans[i] = make([]int, row)
	}
	for i, u := range grid {
		for j, v := range u {
			trans[j][i] = v
		}
	}
	return trans
}

func flip(grid [][]int) [][]int {
	row := len(grid)
	if row == 0 {
		return nil
	}
	col := len(grid[0])
	f := make([][]int, row)
	for i := 0; i < row; i++ {
		f[i] = make([]int, col)
	}
	for i, u := range grid {
		for j, v := range u {
			f[i][col-j-1] = v
		}
	}
	return f
}

func product(selection []int) int {
	p := 1
	for _, v := range selection {
		p *= v
	}
	return p
}

func east(grid [][]int, size int, reducer func([]int) int) []int {
	var results []int
	for _, u := range grid {
		for i := 0; i < len(u)-size+1; i++ {
			s := u[i : i+size]
			results = append(results, reducer(s))
		}
	}
	return results
}

func southeast(grid [][]int, size int, reducer func([]int) int) []int {
	var results []int
	a := len(grid)
	if a == 0 {
		return nil
	}
	b := len(grid[0])
	for i := 0; i < a-size+1; i++ {
		for j := 0; j < b-size+1; j++ {
			s := make([]int, size)
			for k := 0; k < size; k++ {
				s[k] = grid[i+k][j+k]
			}
			results = append(results, reducer(s))
		}
	}
	return results
}

func extend(slices ...[]int) []int {
	var ret []int
	for _, v := range slices {
		ret = append(ret, v...)
	}
	return ret
}

func max(a []int) int {
	var m int
	for i, v := range a {
		if i == 0 {
			m = v
		} else if v > m {
			m = v
		}
	}
	return m
}

func main() {
	g := read("./input.txt")
	w := 4
	p := product
	s := extend(east(g, w, p), southeast(g, w, p), east(transpose(g), w, p), southeast(flip(g), w, p))
	fmt.Println(max(s))
}

// Greatest product of n-adjacent numbers in direction: up, down, left, right, or diagonally.
