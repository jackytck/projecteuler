package main

import (
	"fmt"

	lane "gopkg.in/oleiade/lane.v1"

	"github.com/jackytck/projecteuler/tools"
)

// Node represents a N-gon number.
type Node struct {
	fig   int
	num   int
	left  int
	right int
	child []*Node
}

// Chain represents a chain of Nodes.
type Chain struct {
	nodes []*Node
}

func add(a []int, x int) []int {
	if x < 1000 || x > 9999 {
		return a
	}
	return append(a, x)
}

func initNumber() ([]int, []int, []int, []int, []int, []int) {
	var p3, p4, p5, p6, p7, p8 []int
	n := 1
	for {
		x3 := tools.TriangleNumber(n)
		p3 = add(p3, x3)
		p4 = add(p4, tools.SquareNumber(n))
		p5 = add(p5, tools.PentagonNumber(n))
		p6 = add(p6, tools.HexagonNumber(n))
		p7 = add(p7, tools.HeptagonalNumber(n))
		p8 = add(p8, tools.OctagonalNumber(n))
		n++
		if x3 > 9999 {
			break
		}
	}
	return p3, p4, p5, p6, p7, p8
}

func initNodes(p []int, c int, ns []Node) []Node {
	for _, p := range p {
		d := tools.Digits(p)
		left := tools.JoinInts(d[:2])
		right := tools.JoinInts(d[2:])
		n := Node{c, p, left, right, nil}
		ns = append(ns, n)
	}
	return ns
}

func initEdges(nodes []Node) {
	for i, a := range nodes {
		for j, b := range nodes {
			if i == j || a.fig == b.fig || a.right != b.left {
				continue
			}
			nodes[i].child = append(nodes[i].child, &nodes[j])
		}
	}
}

func bfs(root *Node) (Chain, bool) {
	var found bool
	var chain Chain
	q := lane.NewQueue()
	child := []*Node{root}
	rootChain := Chain{child}
	q.Enqueue(rootChain)
	for !q.Empty() {
		f := q.Dequeue().(Chain)
		if len(f.nodes) == 6 {
			if f.nodes[5].right != f.nodes[0].left {
				continue
			}
			found = true
			chain = f
			break
		}

		s := len(f.nodes)
		existedFig := make(map[int]bool)
		for _, v := range f.nodes {
			existedFig[v.fig] = true
		}

		for _, c := range f.nodes[s-1].child {
			if existedFig[c.fig] {
				continue
			}
			var child []*Node
			for _, fc := range f.nodes {
				child = append(child, fc)
			}
			child = append(child, c)
			next := Chain{child}
			q.Enqueue(next)
		}
	}
	return chain, found
}

func solve() int {
	var ans int
	p3, p4, p5, p6, p7, p8 := initNumber()

	var nodes []Node
	nodes = initNodes(p3, 3, nodes)
	nodes = initNodes(p4, 4, nodes)
	nodes = initNodes(p5, 5, nodes)
	nodes = initNodes(p6, 6, nodes)
	nodes = initNodes(p7, 7, nodes)
	nodes = initNodes(p8, 8, nodes)

	initEdges(nodes)

	for i, v := range nodes {
		if len(v.child) == 0 {
			continue
		}
		if chain, found := bfs(&nodes[i]); found {
			for _, n := range chain.nodes {
				// fmt.Println(n.fig, n.num)
				ans += n.num
			}
			break
		}
	}
	return ans
}

func main() {
	fmt.Println(solve())
}

// Find the sum of the only ordered set of six cyclic 4-digit numbers for which
// each polygonal type: triangle, square, pentagonal, hexagonal, heptagonal,
// and octagonal, is represented by a different number in the set.
// Note:
// Although the total permutation space is huge, the number of edges of each
// node is small, so it could be bfs efficiently.
