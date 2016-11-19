package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strconv"
	"strings"
)

func read(path string) []*big.Int {
	var nums []*big.Int
	lines, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	ss := strings.Split(string(lines), "\n")
	row := len(ss) - 1
	if row == 0 {
		log.Fatal()
	}
	for _, v := range ss {
		if v == "" {
			continue
		}
		n := new(big.Int)
		fmt.Sscan(v, n)
		nums = append(nums, n)
	}
	return nums
}

func solve(input string) int {
	nums := read(input)
	sum := big.NewInt(0)
	for _, n := range nums {
		sum.Add(sum, n)
	}
	ans, _ := strconv.Atoi(sum.String()[:10])
	return ans
}

func main() {
	fmt.Println(solve("./input.txt"))
}

// Find the first 10 digits of the sum of 100 50-digits numbers.
