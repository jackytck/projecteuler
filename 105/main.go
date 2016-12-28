package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/jackytck/projecteuler/tools"
)

func load(input string) [][]int {
	var all [][]int
	for _, s := range tools.ReadFile(input) {
		var set []int
		for _, x := range strings.Split(s, ",") {
			y, _ := strconv.Atoi(x)
			set = append(set, y)
		}
		all = append(all, set)
	}
	return all
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

func specialSumSet(s []int) bool {
	return rule2(s) && rule1(s)
}

// producer
func gen(sets [][]int) <-chan []int {
	out := make(chan []int)
	go func() {
		for _, s := range sets {
			out <- s
		}
		close(out)
	}()
	return out
}

// worker
func work(sets <-chan []int) <-chan int {
	out := make(chan int)
	go func() {
		for s := range sets {
			if specialSumSet(s) {
				out <- tools.Sum(s...)
			}
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func solve(input string) int {
	var cnt int
	sets := load(input)
	g := gen(sets)
	con := 2
	var w []<-chan int
	for i := 0; i < con; i++ {
		w = append(w, work(g))
	}
	for c := range merge(w...) {
		cnt += c
	}
	return cnt
}

func main() {
	fmt.Println(solve("p105_sets.txt"))
}

// Using sets.txt, a 4K text file with one-hundred sets containing seven to
// twelve elements, identify all the special sum sets, A1, A2, ..., Ak, and find
// the value of S(A1) + S(A2) + ... + S(Ak).
// Straight forward ways to check rule 1 and 2.
// Computation is speeded up with concurrency using the pattern described in:
// https://blog.golang.org/pipelines
