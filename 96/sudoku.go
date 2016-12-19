package main

import (
	"strconv"

	lane "gopkg.in/oleiade/lane.v1"

	"github.com/jackytck/projecteuler/tools"
)

// Sudoku represents a Sudoku board.
type Sudoku struct {
	data [][]int
}

// Init initializes the data entries of Sudoki.
func (s *Sudoku) Init(str []string) {
	s.data = make([][]int, 9)
	for i := 0; i < 9; i++ {
		s.data[i] = make([]int, 9)
	}
	for i, r := range str {
		for j, c := range r {
			a, _ := strconv.Atoi(string(c))
			s.data[i][j] = a
		}
	}
}

// Row gives the entries of the i-th row.
func (s *Sudoku) Row(i int) []int {
	r := make([]int, 9)
	if i < 0 || i >= 9 {
		return r
	}
	for j, v := range s.data[i] {
		r[j] = v
	}
	return r
}

// Col gives the entries of the i-th column.
func (s *Sudoku) Col(i int) []int {
	r := make([]int, 9)
	if i < 0 || i >= 9 {
		return r
	}
	for j := 0; j < 9; j++ {
		r[j] = s.data[j][i]
	}
	return r
}

// Neg gives the 3x3 neighbours (excluding itself) of a given cell.
func (s *Sudoku) Neg(i, j int) []int {
	var r []int
	if i < 0 || i >= 9 || j < 0 || j >= 9 {
		return r
	}
	x, y := i/3, j/3
	for u := 3 * x; u < 3*x+3; u++ {
		for v := 3 * y; v < 3*y+3; v++ {
			if u == i && v == j {
				continue
			}
			r = append(r, s.data[u][v])
		}
	}
	return r
}

// Candidate gives the possible entries of the current location.
func (s *Sudoku) Candidate(i, j int) []int {
	var c, e []int
	if i < 0 || i >= 9 || j < 0 || j >= 9 || s.data[i][j] != 0 {
		return c
	}
	e = append(e, s.Row(i)...)
	e = append(e, s.Col(j)...)
	e = append(e, s.Neg(i, j)...)
	e = tools.SetInt(e)
	for k := 1; k <= 9; k++ {
		if tools.IncludesInt(e, k) {
			continue
		}
		c = append(c, k)
	}
	return c
}

func (s *Sudoku) isFull() bool {
	ret := true
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.data[i][j] == 0 {
				ret = false
				break
			}
		}
	}
	return ret
}

func (s *Sudoku) leastCandidate() (int, int, bool) {
	var found bool
	var minI, minJ, minSize int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if c := s.Candidate(i, j); s.data[i][j] == 0 && (minSize == 0 || len(c) < minSize) {
				found = true
				minI, minJ, minSize = i, j, len(c)
			}
		}
	}
	return minI, minJ, found
}

func (s *Sudoku) copy() Sudoku {
	su := Sudoku{}
	su.data = make([][]int, 9)
	for i := 0; i < 9; i++ {
		su.data[i] = append(su.data[i], s.data[i]...)
	}
	return su
}

func (s *Sudoku) setFrom(other *Sudoku) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			s.data[i][j] = other.data[i][j]
		}
	}
}

// Solve solves the sudoku in one possible way.
func (s *Sudoku) Solve() {
	root := s.copy()
	// bfs
	q := lane.NewQueue()
	q.Enqueue(root)
	for !q.Empty() {
		f := q.Dequeue().(Sudoku)
		if f.isFull() {
			s.setFrom(&f)
			break
		}
		x, y, ok := f.leastCandidate()
		if !ok {
			continue
		}
		for _, c := range f.Candidate(x, y) {
			child := f.copy()
			child.data[x][y] = c
			q.Enqueue(child)
		}
	}
}

// Hash gives the 3-digit numbers found in the top left corner.
func (s *Sudoku) Hash() int {
	return tools.JoinInts(s.data[0][0:3])
}

func (s Sudoku) String() string {
	var str string
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			str += "-------------------------\n"
		}
		for j := 0; j < 8; j++ {
			if j%3 == 0 {
				str += "| "
			}
			str += strconv.Itoa(s.data[i][j]) + " "
		}
		str += strconv.Itoa(s.data[i][8]) + " |\n"
	}
	str += "-------------------------"
	return str
}
