package main

import "fmt"

func brute() int {
	var cnt int
	// 100, 50, 20, 10, 5, 2, 1
	for a := 0; a <= 2; a++ {
		for b := 0; b <= 4; b++ {
			for c := 0; c <= 10; c++ {
				for d := 0; d <= 20; d++ {
					for e := 0; e <= 40; e++ {
						for f := 0; f <= 100; f++ {
							for g := 0; g <= 200; g++ {
								if p := a*100 + b*50 + c*20 + d*10 + e*5 + f*2 + g; p == 200 {
									cnt++
								}
							}
						}
					}
				}
			}
		}
	}
	return cnt + 1
}

func dp() int {
	coins := [8]int{1, 2, 5, 10, 20, 50, 100, 200}
	table := make(map[int]int)
	table[0] = 1
	for _, c := range coins {
		for i := c; i <= 200; i++ {
			table[i] += table[i-c]
		}
	}
	return table[200]
}

func main() {
	// fmt.Println(brute())
	fmt.Println(dp())
}

// Number of ways to add up to 200 by 1, 2, 5, 10, 20, 50, 100 and 200.
// Note:
// If I know the number of ways to add up to 1 to 200, by using coins 1, 2, 5 only.
// Then finding the number of ways to add up to 1 to 200 by an additional coin
// e.g. 10 would be very easy. Key insight is to add new coins (not the sum of value)
// incrementally.
