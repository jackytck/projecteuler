package main

import (
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"

	"github.com/jackytck/projecteuler/tools"
)

func cardRank(card string) int {
	order := map[string]int{
		"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8,
		"9": 9, "T": 10, "J": 11, "Q": 12, "K": 13, "A": 14,
		"S": 4, "H": 3, "C": 2, "D": 1}
	num := card[:len(card)-1]
	// suit := string(card[len(card)-1])
	// return order[num]*10 + order[suit]
	return order[num]
}

func parseHands(s string) ([]string, []string) {
	hands := strings.Split(s, " ")
	return hands[:5], hands[5:]
}

func pow(a, b int) int {
	return int(tools.Exp(a, b).Int64())
}

func countDigit(hand []string) map[string]int {
	c := make(map[string]int)
	for _, v := range hand {
		c[string(v[0])]++
	}
	return c
}

func countSuit(hand []string) map[string]int {
	c := make(map[string]int)
	for _, v := range hand {
		c[string(v[1])]++
	}
	return c
}

func highCard(hand []string) int {
	var r int
	var a []int
	for _, c := range hand {
		a = append(a, cardRank(c))
	}
	sort.Ints(a)
	for i, c := range a {
		r += c * pow(15, i)
	}
	return r
}

func pairs(hand []string) int {
	var score int
	counter := countDigit(hand)
	var pairs []int
	var threeKind bool
	for k, v := range counter {
		if v == 2 {
			pairs = append(pairs, cardRank(k+" "))
		} else if v == 3 {
			threeKind = true
		}
	}
	sort.Ints(pairs)
	switch len(pairs) {
	case 1:
		score = pow(15, 7) + pow(15, 6)*pairs[0]
		if threeKind {
			score = 0
		}
	case 2:
		score = pow(15, 7)*2 + pow(15, 6)*pairs[1] + pow(15, 5)*pairs[0]
	}
	return score
}

func threeKind(hand []string) int {
	var score int
	counter := countDigit(hand)
	var pairs bool
	for k, v := range counter {
		if v == 3 {
			score = pow(15, 7)*3 + pow(15, 6)*cardRank(k+" ")
		} else if v == 2 {
			pairs = true
		}
	}
	if pairs {
		score = 0
	}
	return score
}

func isStraight(hand []string) bool {
	var digits []int
	for _, v := range hand {
		digits = append(digits, cardRank(v))
	}
	sort.Ints(digits)
	m := digits[0]
	for i := range digits {
		digits[i] -= m
	}
	return tools.JoinInts(digits) == 1234 ||
		reflect.DeepEqual(digits, []int{0, 1, 2, 3, 12})
}

func isFlush(hand []string) bool {
	var flush bool
	counter := countSuit(hand)
	for _, v := range counter {
		if v == 5 {
			flush = true
			break
		}
	}
	return flush
}

func straight(hand []string) int {
	var score int
	if isStraight(hand) && !isFlush(hand) {
		score = pow(15, 7) * 4
	}
	return score
}

func flush(hand []string) int {
	var score int
	if isFlush(hand) && !isStraight(hand) {
		score = pow(15, 7) * 5
	}
	return score
}

func fullHouse(hand []string) int {
	var score int
	counter := countDigit(hand)
	var pairs bool
	var kind bool
	var value int
	for k, v := range counter {
		if v == 3 {
			kind = true
			value = cardRank(k + " ")
		} else if v == 2 {
			pairs = true
		}
	}
	if kind && pairs {
		score = pow(15, 7)*6 + pow(15, 6)*value
	}
	return score
}

func fourKind(hand []string) int {
	var score int
	counter := countDigit(hand)
	for k, v := range counter {
		if v == 4 {
			score = pow(15, 7)*7 + pow(15, 6)*cardRank(k+" ")
			break
		}
	}
	return score
}

func straightFlush(hand []string) int {
	var score int
	if isStraight(hand) && isFlush(hand) {
		score = pow(15, 7) * 8
	}
	return score
}

func royalFlush(hand []string) int {
	score := straightFlush(hand)
	var ten bool
	var ace bool
	for _, v := range hand {
		if v == "T" {
			ten = true
		} else if v == "A" {
			ace = true
		}
	}
	if score > 0 && ten && ace {
		score = pow(15, 7) * 9
	}
	return score
}

func handScore(hand []string) int {
	return highCard(hand) + pairs(hand) + threeKind(hand) + straight(hand) +
		flush(hand) + fullHouse(hand) + fourKind(hand) + straightFlush(hand) +
		royalFlush(hand)
}

func compareHands(p1, p2 []string) bool {
	return handScore(p1) > handScore(p2)
}

func test() {
	s1 := "5H 5C 6S 7S KD 2C 3S 8S 8D TD"
	s2 := "5D 8C 9S JS AC 2C 5C 7D 8S QH"
	s3 := "2D 9C AS AH AC 3D 6D 7D TD QD"
	s4 := "4D 6S 9H QH QC 3D 6D 7H QD QS"
	s5 := "2H 2D 4C 4D 4S 3C 3D 3S 9S 9D"
	tests := []string{s1, s2, s3, s4, s5}
	for i, v := range tests {
		if compareHands(parseHands(v)) {
			fmt.Println(i+1, "Player 1")
		} else {
			fmt.Println(i+1, "Player 2")
		}
	}
}

func solve(input string) int {
	var win int
	inputs, err := tools.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	for _, h := range inputs {
		if compareHands(parseHands(h)) {
			win++
		}
	}
	fmt.Println(win)
	return win
}

func main() {
	fmt.Println(solve("./p054_poker.txt"))
}

// How many hands does Player 1 win?
