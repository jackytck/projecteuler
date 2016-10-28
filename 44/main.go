package main

import (
	"fmt"

	"github.com/jackytck/projecteuler/tools"
)

func main() {
	var minDiff int
	var found bool
	var definite bool
	j := 1
	for !definite {
		j++
		jn := tools.PentagonNumber(j)
		for k := j - 1; k > 0; k-- {
			kn := tools.PentagonNumber(k)
			if found && k == j-1 && jn-kn > minDiff {
				definite = true
				break
			}
			sum := jn + kn
			diff := jn - kn
			if found && diff > minDiff {
				continue
			}
			if tools.IsPentagonNumber(sum) && tools.IsPentagonNumber(diff) {
				if !found || diff < minDiff {
					minDiff = diff
				}
				found = true
				break
			}
		}
	}
	fmt.Println(minDiff)
}

// Find the pair of pentagonal numbers, Pj and Pk, for which their sum and
// difference are pentagonal and D = |Pk − Pj| is minimised.
// Note:
// After the best possible answer is found, it is compared with P(j)-P(j-1) for
// ever increasing j, until a bigger difference is found. At this point,
// the definitive answer is found.
