package main

import (
	"fmt"
	"math/rand"
	"sort"
)

const (
	gogo = iota
	a1
	cc1
	a2
	t1
	r1
	b1
	ch1
	b2
	b3
	jail
	c1
	u1
	c2
	c3
	r2
	d1
	cc2
	d2
	d3
	fp
	e1
	ch2
	e2
	e3
	r3
	f1
	f2
	u2
	f3
	g2j
	g1
	g2
	cc3
	g3
	r4
	ch3
	h1
	t2
	h2
)

func roll(side int) (int, bool) {
	a := rand.Intn(side) + 1
	b := rand.Intn(side) + 1
	return a + b, a == b
}

func chest(pos int) int {
	d := rand.Intn(16) + 1
	if d == 1 {
		return gogo
	} else if d == 2 {
		return jail
	}
	return pos
}

func nextRail(pos int) int {
	switch {
	case pos < r1:
		return r1
	case pos < r2:
		return r2
	case pos < r3:
		return r3
	case pos < r4:
		return r4
	}
	return r1
}

func nextUtility(pos int) int {
	switch {
	case pos < u1:
		return u1
	case pos < u2:
		return u2
	}
	return u1
}

func chance(pos int) int {
	d := rand.Intn(16) + 1
	switch d {
	case 1:
		return gogo
	case 2:
		return jail
	case 3:
		return c1
	case 4:
		return e3
	case 5:
		return h2
	case 6:
		return r1
	case 7:
		return nextRail(pos)
	case 8:
		return nextRail(pos)
	case 9:
		return nextUtility(pos)
	case 10:
		return pos - 3
	}
	return pos
}

func isChest(pos int) bool {
	if pos == cc1 || pos == cc2 || pos == cc3 {
		return true
	}
	return false
}

func isChance(pos int) bool {
	if pos == ch1 || pos == ch2 || pos == ch3 {
		return true
	}
	return false
}

func next(pos, double, side int) (int, int) {
	advance, isDouble := roll(side)
	if isDouble {
		double++
	} else {
		double = 0
	}
	pos = (pos + advance) % 40
	if double == 3 || pos == g2j {
		return jail, double % 3
	}
	if isChest(pos) {
		return chest(pos), double
	}
	if isChance(pos) {
		return chance(pos), double
	}
	return pos, double
}

type slice struct {
	sort.IntSlice
	idx []int
}

func (s slice) Swap(i, j int) {
	s.IntSlice.Swap(i, j)
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

func (s slice) Less(i, j int) bool {
	return !s.IntSlice.Less(i, j)
}

func newSlice(n ...int) *slice {
	s := &slice{IntSlice: sort.IntSlice(n), idx: make([]int, len(n))}
	for i := range s.idx {
		s.idx[i] = i
	}
	return s
}

func monteCarlo(trial, side int) string {
	boards := make([]int, 40)
	pos := gogo
	var double int
	for i := 0; i < trial; i++ {
		pos, double = next(pos, double, side)
		boards[pos]++
	}
	s := newSlice(boards...)
	sort.Sort(s)
	modal := fmt.Sprintf("%02d%02d%02d", s.idx[0], s.idx[1], s.idx[2])
	return modal
}

func main() {
	fmt.Println(monteCarlo(10000000, 6))
	fmt.Println(monteCarlo(10000000, 4))
}

// Find the three most popular squares of a Monopoly board game using two
// 4-sided dice.
// Note:
// Method of Monte Carlo is used.
