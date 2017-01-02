package main

import (
	"fmt"
	"math"

	"github.com/jackytck/projecteuler/tools"
)

var log = math.Log10

func m(a, b int) int {
	return tools.MinInt(a, b)
}

func over(es ...int) bool {
	d := 1
	for _, v := range es {
		if v == 0 {
			continue
		}
		d *= 2*v + 1
	}
	return d > 8000000
}

func solve() int64 {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43}
	limit := int(tools.Exp(3, 14).Int64())
	logLimit := log(float64(limit))
	logP := make([]float64, len(primes))
	bs := make([]int, len(primes))
	for i, p := range primes {
		bs[i] = int(logLimit / log(float64(p)))
		logP[i] = log(float64(p))
	}

	mul := func(es ...int) float64 {
		var p float64
		for i, v := range es {
			p += float64(v) * logP[i]
		}
		return p
	}

	var logAns float64
	var minE []int
	for e1 := 2; e1 < bs[0]; e1++ {
		for e2 := 0; e2 <= m(e1, bs[1]); e2++ {
			for e3 := 0; e3 <= m(e2, bs[2]); e3++ {
				for e4 := 0; e4 <= m(e3, bs[3]); e4++ {
					for e5 := 0; e5 <= m(e4, bs[4]); e5++ {
						for e6 := 0; e6 <= m(e5, bs[5]); e6++ {
							for e7 := 0; e7 <= m(e6, bs[6]); e7++ {
								for e8 := 0; e8 <= m(e7, bs[7]); e8++ {
									for e9 := 0; e9 <= m(e8, bs[8]); e9++ {
										for e10 := 0; e10 <= m(e9, bs[9]); e10++ {
											for e11 := 0; e11 <= m(e10, bs[10]); e11++ {
												for e12 := 0; e12 <= m(e11, bs[11]); e12++ {
													for e13 := 0; e13 <= m(e12, bs[12]); e13++ {
														for e14 := 0; e14 <= m(e13, bs[13]); e14++ {
															es := []int{e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13, e14}
															if over(es...) {
																if p := mul(es...); logAns == 0 || p < logAns {
																	logAns = p
																	minE = es
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	var ans int64 = 1
	for i, v := range minE {
		z := tools.Exp(primes[i], v)
		ans *= z.Int64()
	}
	return ans
}

func main() {
	fmt.Println(solve())
}

// In the following equation x, y, and n are positive integers.
// 1/x + 1/y = 1/n
// What is the least value of n for which the number of distinct solutions
// exceeds four million?
// Note:
// Instead of doing the prime factorizations for each number as in problem 108.
// Just exhaustively try out all the exponents of primes and keep track of the
// minimum.
