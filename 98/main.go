package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/jackytck/projecteuler/tools"
)

// Pair represents an anagram word pair.
type Pair struct {
	a, b string
}

func hash(w string) string {
	t := strings.Split(w, "")
	sort.Strings(t)
	return strings.Join(t, "")
}

func posiblePairs(words []string) ([]Pair, int, int) {
	var pairs []Pair
	var minW, maxW int
	dict := make(map[string][]string)
	for _, w := range words {
		h := hash(w)
		if _, ok := dict[h]; ok {
			dict[h] = append(dict[h], w)
		} else {
			dict[h] = []string{w}
		}
	}
	for k, v := range dict {
		if len(v) == 1 {
			continue
		}
		for c := range tools.CombIndex(len(v), 2) {
			p := Pair{v[c[0]], v[c[1]]}
			if p.a != p.b {
				pairs = append(pairs, p)
			}
		}
		wl := len(k)
		if minW == 0 {
			minW = wl
		}
		minW = tools.MinInt(minW, wl)
		maxW = tools.MaxInt(maxW, wl)
	}
	return pairs, minW, maxW
}

func genSquares(minDigit, maxDigit int) map[int][]int {
	m := make(map[int][]int)
	start := int(math.Ceil(math.Sqrt((math.Pow10(minDigit - 1)))))
	end := int(math.Pow10(maxDigit))
	var s int
	for i := start; s < end; i++ {
		s = i * i
		if s > end {
			break
		}
		k := int(math.Log10(float64(s))) + 1
		if _, ok := m[k]; ok {
			m[k] = append(m[k], s)
		} else {
			m[k] = []int{s}
		}
	}
	return m
}

func isAnagramicSquare(p Pair, sq int) (bool, int) {
	// map first pair to square number
	m := make(map[rune]int)
	digits := tools.Digits(sq)
	var i int
	for _, v := range p.a {
		if _, ok := m[v]; !ok {
			m[v] = digits[i]
			i++
		}
	}

	// check if same rune matches to the same int
	visited := make(map[int]bool)
	for _, v := range m {
		if _, ok := visited[v]; ok {
			return false, 0
		}
		visited[v] = true
	}

	// final transformed number
	var first []int
	for _, v := range p.a {
		first = append(first, m[v])
	}
	var second []int
	for _, v := range p.b {
		second = append(second, m[v])
	}
	a, b := tools.JoinInts(first), tools.JoinInts(second)

	// no leading zero and is a square number
	return second[0] != 0 && tools.IsSquareNumber(a) && tools.IsSquareNumber(b),
		tools.MaxInt(a, b)
}

func solve(input string) int {
	var maxSq int
	words := tools.ReadWords(input)
	pairs, min, max := posiblePairs(words)
	squares := genSquares(min, max)
	for _, p := range pairs {
		for _, sq := range squares[tools.MinInt(len(p.a), len(p.b))] {
			if valid, s := isAnagramicSquare(p, sq); valid {
				maxSq = tools.MaxInt(maxSq, s)
			}
		}
	}
	return maxSq
}

func main() {
	fmt.Println(solve("p098_words.txt"))
}

// Using p098_words.txt, find all the square anagram word pairs. What is the
// largest square number formed by any member of such a pair?
// Note:
// First find all the possible pairs, then pre-generate all required square
// numbers. Finally match each pair with each possible square and check its
// validity.
