package main

import "fmt"

// took 3m9.427s
func dp() int {
	divisor := 1000000
	table := make(map[int]int)
	table[0] = 1
	for c := 1; c < 60000; c++ {
		for i := c; i <= 60000; i++ {
			table[i] += table[i-c] % divisor
		}
		if table[c]%divisor == 0 {
			return c
		}
	}
	return 0
}

func p(n, mod int) []int {
	if n < 0 {
		return nil
	}
	if n == 0 {
		return []int{1}
	}
	if n == 1 {
		return []int{0, 1}
	}
	pentagon := make(map[int]int)
	for k := 1; true; k++ {
		pentagon[k] = k * (3*k - 1) / 2
		pentagon[-k] = k * (3*k + 1) / 2
		if pentagon[k] > n && pentagon[-k] > n {
			break
		}
	}
	toMod := false
	if mod > 0 {
		toMod = true
	}
	table := make([]int, n+1)
	table[0] = 1
	table[1] = 1
	for i := 2; i <= n; i++ {
		s := -1
		for k := 1; true; k++ {
			g1 := pentagon[k]
			g2 := pentagon[-k]
			s *= -1
			if g1 <= i {
				table[i] += s * table[i-g1]
			}
			if g2 <= i {
				table[i] += s * table[i-g2]
			}
			if toMod {
				table[i] = table[i] % mod
			}
			if g1 > i && g2 > i {
				break
			}
		}
	}
	return table
}

// took 0m0.458s
func solve() int {
	table := p(60000, 1000000)
	for i, v := range table {
		if v == 0 {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(solve())
}

// Find the least value of n for which p(n) is divisible by one million.
// Note:
// This is in fact the same DP as in problem 31. But the table is too large
// for the 2D DP. So the recurrence formula is used instead:
// https://en.wikipedia.org/wiki/Partition_(number_theory)#Recurrence_formula
