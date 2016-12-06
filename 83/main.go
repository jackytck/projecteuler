package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
	lane "gopkg.in/oleiade/lane.v1"
)

var w, h int

type node struct {
	i, j, v int
}

func (n *node) valid() bool {
	if n.i >= 0 && n.i < h && n.j >= 0 && n.j < w {
		return true
	}
	return false
}

func (n *node) child() []node {
	var c []node
	left := node{n.i, n.j - 1, 0}
	right := node{n.i, n.j + 1, 0}
	up := node{n.i - 1, n.j, 0}
	down := node{n.i + 1, n.j, 0}
	if left.valid() {
		c = append(c, left)
	}
	if right.valid() {
		c = append(c, right)
	}
	if up.valid() {
		c = append(c, up)
	}
	if down.valid() {
		c = append(c, down)
	}
	return c
}

func (n *node) hash() int {
	return n.i*h + n.j
}

func solve(path string) int {
	m := tools.ReadMatrix(path)
	h = len(m)
	w = len(m[0])
	root := node{0, 0, m[0][0]}

	// Dijkstra
	var sum int
	visit := make(map[int]bool)
	pq := lane.NewPQueue(lane.MINPQ)
	pq.Push(root, root.v)
	for !pq.Empty() {
		t, v := pq.Pop()
		tn := t.(node)
		hash := tn.hash()
		if tn.i == h-1 && tn.j == w-1 {
			sum = v
			break
		}
		if _, ok := visit[hash]; ok {
			continue
		}
		visit[hash] = true
		for _, c := range tn.child() {
			c.v = m[c.i][c.j] + v
			pq.Push(c, c.v)
		}
	}
	return sum
}

func main() {
	fmt.Println(solve("./p083_matrix_small.txt"))
	fmt.Println(solve("./p083_matrix.txt"))
}

// Find the minimal path sum of a matrix from the top left to the bottom right
// by moving left, right, up, and down.
// Note:
// Classic Dijkstra. Same problem as p82. The only difference is now the node
// could go left.
