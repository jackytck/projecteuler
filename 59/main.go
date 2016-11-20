package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/jackytck/projecteuler/tools"
)

func read(path string) []int {
	var ret []int
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	slice := strings.Split(strings.TrimSpace(string(data)), ",")
	for _, v := range slice {
		i, _ := strconv.Atoi(v)
		ret = append(ret, i)
	}
	return ret
}

func keys() chan string {
	c := make(chan string)
	go func() {
		for i := 'a'; i <= 'z'; i++ {
			for j := 'a'; j <= 'z'; j++ {
				for k := 'a'; k <= 'z'; k++ {
					c <- string(i) + string(j) + string(k)
				}
			}
		}
		close(c)
	}()
	return c
}

func decrypt(text []int, key string) (string, []int) {
	var decoded []string
	var ascii []int
	for i, v := range text {
		k := key[i%3]
		w := v ^ int(k)
		decoded = append(decoded, string(w))
		ascii = append(ascii, w)
	}
	return strings.Join(decoded, ""), ascii
}

func score(text string) int {
	var cnt int
	for _, v := range text {
		if v == ' ' {
			cnt++
		}
	}
	return cnt
}

func solve(input string) int {
	// var key, raw string
	var maxScore int
	var ascii []int
	encrypted := read(input)
	for k := range keys() {
		trial, num := decrypt(encrypted, k)
		if s := score(trial); s > maxScore {
			maxScore = s
			// key = k
			// raw = trial
			ascii = num
		}
	}
	// fmt.Println(raw)
	// fmt.Println(key, maxScore)
	return tools.Sum(ascii...)
}

func main() {
	fmt.Println(solve("./p059_cipher.txt"))
}

// Decrypt the message and find the sum of the ASCII values in the original text.
// Note:
// All possible keys are tested. The best one is picked based on the number of
// occurrences of the space character in the deciphered text.
