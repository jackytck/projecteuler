package main

import (
	"fmt"
	"log"

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
	right := node{n.i, n.j + 1, 0}
	up := node{n.i - 1, n.j, 0}
	down := node{n.i + 1, n.j, 0}
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
	m, err := tools.ReadMatrix(path)
	if err != nil {
		log.Fatal(err)
	}
	h = len(m)
	w = len(m[0])
	var roots []node
	for i := 0; i < h; i++ {
		roots = append(roots, node{i, 0, m[i][0]})
	}

	// Dijkstra
	var minSum int
	for _, r := range roots {
		var sum int
		visit := make(map[int]bool)
		pq := lane.NewPQueue(lane.MINPQ)
		pq.Push(r, r.v)
		for !pq.Empty() {
			t, v := pq.Pop()
			tn := t.(node)
			hash := tn.hash()
			if tn.j == w-1 {
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
		if minSum == 0 || sum < minSum {
			minSum = sum
		}
	}
	return minSum
}

func main() {
	fmt.Println(solve("./p082_matrix_small.txt"))
	fmt.Println(solve("./p082_matrix.txt"))
}

// Find the minimal path sum of a matrix from the left column to the right
// column.
// Note:
// Classic Dijkstra.
