package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	lane "gopkg.in/oleiade/lane.v1"

	"github.com/jackytck/projecteuler/tools"
)

// Node represents a set of size n.
type Node struct {
	set []int
}

func (n *Node) copy() Node {
	s := make([]int, len(n.set))
	copy(s, n.set)
	return Node{s}
}

func (n *Node) child() []Node {
	m := make(map[int]bool)
	for _, v := range n.set {
		m[v] = true
	}
	var ret []Node
	for i, v := range n.set {
		k := v + 1
		if _, ok := m[k]; !ok {
			c := n.copy()
			c.set[i] = k
			ret = append(ret, c)
		}
	}
	return ret
}

func (n *Node) isSpecialSumSet() bool {
	return rule2(n.set) && rule1(n.set)
}

func (n *Node) hash() string {
	s := make([]string, len(n.set))
	for i, v := range n.set {
		s[i] = strconv.Itoa(v)
	}
	return strings.Join(s, ",")
}

func powerSet(s []int) [][]int {
	var p [][]int
	for ids := range tools.CartProduct(2, len(s)) {
		var x []int
		for i, v := range ids {
			if v == 1 {
				x = append(x, s[i])
			}
		}
		p = append(p, x)
	}
	return p
}

func sum(s [][]int) []int {
	r := make([]int, len(s))
	for i, v := range s {
		r[i] = tools.Sum(v...)
	}
	return r
}

func distjoint(i, j int) bool {
	return i != 0 && i&j == 0
}

func rule1(s []int) bool {
	pass := true
	p := sum(powerSet(s))
	for pair := range tools.CombIndex(len(p), 2) {
		i, j := pair[0], pair[1]
		if distjoint(i, j) && p[i] == p[j] {
			pass = false
			break
		}
	}
	return pass
}

func rule2(s []int) bool {
	pass := true
	size := len(s)
	a := make([]int, size)
	copy(a, s)
	sort.Ints(a)
	for i := 2; i+i-1 <= size; i++ {
		sa := tools.Sum(a[:i]...)
		sb := tools.Sum(a[size+1-i:]...)
		if sa <= sb {
			pass = false
			break
		}
	}
	return pass
}

func solve() string {
	var ans string
	r := Node{[]int{17, 28, 35, 36, 37, 39, 42}}
	q := lane.NewQueue()
	visit := make(map[string]bool)
	q.Enqueue(r)
	for !q.Empty() {
		f := q.Dequeue().(Node)
		h := f.hash()
		if visit[h] {
			continue
		}
		visit[h] = true
		if f.isSpecialSumSet() {
			ans = tools.JoinIntsString(f.set...)
			break
		}
		for _, c := range f.child() {
			q.Enqueue(c)
		}
	}
	return ans
}

func main() {
	fmt.Println(solve())
}

// Given that A is an optimum special sum set for n = 7, find its set string.
// Note:
// It is assumed that the approximate solution given by the algorithm described
// in the description is very closed to the optimal solution. So the root of the
// BFS is set as the node where each element is 3 less than the approximated.
// Then incrementally go up to find the first special sum.
