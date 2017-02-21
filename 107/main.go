package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	lane "gopkg.in/oleiade/lane.v1"

	"github.com/jackytck/projecteuler/tools"
)

// Node represents a node in the graph.
type Node struct {
	id   int
	edge []Edge
}

// Edge represents an edge with weight in the graph.
type Edge struct {
	v, w int
}

func loadGraph(input string) ([]Node, int) {
	var total int
	var nodes []Node
	lines, err := tools.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	for i, line := range lines {
		var edges []Edge
		for j, v := range strings.Split(line, ",") {
			if v == "-" {
				continue
			}
			w, _ := strconv.Atoi(v)
			e := Edge{j, w}
			total += w
			edges = append(edges, e)
		}
		n := Node{i, edges}
		nodes = append(nodes, n)
	}
	return nodes, total / 2
}

func solve(input string) int {
	g, total := loadGraph(input)
	r := g[0]
	pq := lane.NewPQueue(lane.MINPQ)
	pq.Push(r, 0)
	visit := make(map[int]bool)
	var mst int
	for !pq.Empty() {
		t, v := pq.Pop()
		tn := t.(Node)
		if visit[tn.id] {
			continue
		}
		mst += v
		visit[tn.id] = true
		for _, c := range tn.edge {
			pq.Push(g[c.v], c.w)
		}
	}
	return total - mst
}

func main() {
	fmt.Println(solve("p107_network.txt"))
}

// Using network.txt, a 6K text file containing a network with forty vertices,
// and given in matrix form, find the maximum saving which can be achieved by
// removing redundant edges whilst ensuring that the network remains connected.
// Note:
// Classic MST using Prim's algorithm.
