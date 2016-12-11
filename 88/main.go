package main

import "fmt"

func store(p, s, c int, dp map[int]int) {
	k := p - s + c
	if v, ok := dp[k]; ok {
		if p < v {
			dp[k] = p
		}
	} else {
		dp[k] = p
	}
}

func solve() int {
	k := 12000
	dp := make(map[int]int)
	up, up2 := k, 2*k
	for a1 := 2; a1 <= up; a1++ {
		p1, s1 := a1, a1
		for a2 := a1; a2 <= up; a2++ {
			p2 := p1 * a2
			if p2 > up2 {
				break
			}
			s2 := s1 + a2
			store(p2, s2, 2, dp)
			for a3 := a2; a3 <= up; a3++ {
				p3 := p2 * a3
				if p3 > up2 {
					break
				}
				s3 := s2 + a3
				store(p3, s3, 3, dp)
				for a4 := a3; a4 <= up; a4++ {
					p4 := p3 * a4
					if p4 > up2 {
						break
					}
					s4 := s3 + a4
					store(p4, s4, 4, dp)
					for a5 := a4; a5 <= up; a5++ {
						p5 := p4 * a5
						if p5 > up2 {
							break
						}
						s5 := s4 + a5
						store(p5, s5, 5, dp)
						for a6 := a5; a6 <= up; a6++ {
							p6 := p5 * a6
							if p6 > up2 {
								break
							}
							s6 := s5 + a6
							store(p6, s6, 6, dp)
							for a7 := a6; a7 <= up; a7++ {
								p7 := p6 * a7
								if p7 > up2 {
									break
								}
								s7 := s6 + a7
								store(p7, s7, 7, dp)
								for a8 := a7; a8 <= up; a8++ {
									p8 := p7 * a8
									if p8 > up2 {
										break
									}
									s8 := s7 + a8
									store(p8, s8, 8, dp)
									for a9 := a8; a9 <= up; a9++ {
										p9 := p8 * a9
										if p9 > up2 {
											break
										}
										s9 := s8 + a9
										store(p9, s9, 9, dp)
										for a10 := a9; a10 <= up; a10++ {
											p10 := p9 * a10
											if p10 > up2 {
												break
											}
											s10 := s9 + a10
											store(p10, s10, 10, dp)
											for a11 := a10; a11 <= up; a11++ {
												p11 := p10 * a11
												if p11 > up2 {
													break
												}
												s11 := s10 + a11
												store(p11, s11, 11, dp)
												for a12 := a11; a12 <= up; a12++ {
													p12 := p11 * a12
													if p12 > up2 {
														break
													}
													s12 := s11 + a12
													store(p12, s12, 12, dp)
													for a13 := a12; a13 <= up; a13++ {
														p13 := p12 * a13
														if p13 > up2 {
															break
														}
														s13 := s12 + a13
														store(p13, s13, 13, dp)
														for a14 := a13; a14 <= up; a14++ {
															p14 := p13 * a14
															if p14 > up2 {
																break
															}
															s14 := s13 + a14
															store(p14, s14, 14, dp)
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
	set := make(map[int]bool)
	for i := 2; i <= k; i++ {
		set[dp[i]] = true
	}
	var cnt int
	for i := range set {
		cnt += i
	}
	return cnt
}

func main() {
	fmt.Println(solve())
}

// A natural number, N, that can be written as the sum and product of a given
// set of at least two natural numbers, {a1, a2, ... , ak} is called a
// product-sum number: N = a1 + a2 + ... + ak = a1 × a2 × ... × ak.
// For a given set of size, k, we shall call the smallest N with this property a
// minimal product-sum number. What is the sum of all the minimal product-sum
// numbers for 2≤k≤12000?
// Note:
// This is difficult as the space is very large. All of the factorizations are
// generated. Only up to 14 factors are needed. If an intermediate product is
// larger than 2k, then we could backtrack, which saves a lot of time, because
// the growth of products are very fast. So the total number of trial is small.
// I prefer simplicity over elegance, so I simply use 14 nested loops to do the
// job.
