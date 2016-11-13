package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func merge(a, b int) (int, int) {
	da := tools.Digits(a)
	db := tools.Digits(b)
	var ab []int
	var ba []int
	ab = append(ab, da...)
	ab = append(ab, db...)
	ba = append(ba, db...)
	ba = append(ba, da...)
	return tools.JoinInts(ab), tools.JoinInts(ba)
}

func isPrime(m map[int]bool, a, bound int) bool {
	if a < bound {
		return m[a]
	}
	test := tools.IsPrime(a)
	if test {
		m[a] = true
		return true
	}
	return false
}

func friends(primes []int, m map[int]bool, pa, start, bound int) map[int]bool {
	ret := make(map[int]bool)
	for i := start; i < len(primes); i++ {
		pb := primes[i]
		if pb < pa {
			continue
		}
		if ma, mb := merge(pa, pb); isPrime(m, ma, bound) && isPrime(m, mb, bound) {
			ret[pb] = true
		}
	}
	return ret
}

func keys(m map[int]bool) []int {
	var ret []int
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}

func main() {
	bound := 30000
	var minSum int
	var found bool
	primes := tools.SievePrime(bound)
	s := len(primes)
	pm := make(map[int]bool)
	for _, p := range primes {
		pm[p] = true
	}

	for a := 1; a < s-4; a++ {
		pa := primes[a]
		fda := friends(primes, pm, pa, a+1, bound)
		if len(fda) < 4 || (found && pa*5 > minSum) {
			continue
		}
		for b := a + 1; b < s-3; b++ {
			pb := primes[b]
			if _, ok := fda[pb]; !ok || (found && pb*4 > minSum) {
				continue
			}
			fdb := friends(keys(fda), pm, pb, 0, bound)
			if len(fdb) < 3 {
				continue
			}
			for c := b + 1; c < s-2; c++ {
				pc := primes[c]
				if _, ok := fdb[pc]; !ok || (found && pc*3 > minSum) {
					continue
				}
				fdc := friends(keys(fdb), pm, pc, 0, bound)
				if len(fdc) < 2 {
					continue
				}
				for d := c + 1; d < s-1; d++ {
					pd := primes[d]
					if _, ok := fdc[pd]; !ok || (found && pd*2 > minSum) {
						continue
					}
					fdd := friends(keys(fdc), pm, pd, 0, bound)
					if len(fdd) == 0 {
						continue
					}
					for e := d + 1; e < s; e++ {
						pe := primes[e]
						if _, ok := fdd[pe]; !ok || (found && pe > minSum) {
							continue
						}
						fmt.Println(pa, pb, pc, pd, pe)
						if sum := tools.Sum(pa, pb, pc, pd, pe); !found || sum < minSum {
							minSum = sum
						}
						found = true
					}
				}
			}
		}
	}
	fmt.Println(minSum)
}

// Find the lowest sum for a set of five primes for which any two primes
// concatenate to produce another prime.
// Note:
// The exploring space is huge. So a lot of back-trackings are used.
// Since the check for concatenated prime is strong, most of the trial are
// skipped around the outer loops, thus making the problem tractable.
