package tools

import (
	"math/big"
	"sort"
	"testing"
)

func TestSum(t *testing.T) {
	cases := []struct {
		in  []int
		out int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{-3, -12, 7}, -8},
		{[]int{2, 3, 5, 7, 11, 13}, 41},
		{[]int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}, 376},
	}
	for _, c := range cases {
		v := Sum(c.in...)
		if v != c.out {
			t.Errorf("Sum: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestDivisors(t *testing.T) {
	cases := []struct {
		in  int
		out []int
	}{
		{1, []int{1}},
		{2, []int{1}},
		{3, []int{1}},
		{4, []int{1, 2}},
		{5, []int{1}},
		{66, []int{1, 2, 3, 6, 11, 22, 33}},
		{71, []int{1}},
		{72, []int{1, 2, 3, 4, 6, 8, 9, 12, 18, 24, 36}},
		{180, []int{1, 2, 3, 4, 5, 6, 9, 10, 12, 15, 18, 20, 30, 36, 45, 60, 90}},
		{576, []int{1, 2, 3, 4, 6, 8, 9, 12, 16, 18, 24, 32, 36, 48, 64, 72, 96, 144, 192, 288}},
		{972, []int{1, 2, 3, 4, 6, 9, 12, 18, 27, 36, 54, 81, 108, 162, 243, 324, 486}},
	}
	for _, c := range cases {
		d := Divisors(c.in)
		sort.Ints(d)
		if len(d) != len(c.out) {
			t.Errorf("Divisors: len(%v)\tExpected: len(%v)", len(d), len(c.out))
		}
		for i, v := range d {
			if v != c.out[i] {
				t.Errorf("Divisors: %v\tExpected: %v", v, c.out[i])
			}
		}
	}
}

func TestSumDivisors(t *testing.T) {
	cases := []struct {
		in  int
		out int
	}{
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 3},
		{5, 1},
		{66, 78},
		{71, 1},
		{72, 123},
		{180, 366},
		{576, 1075},
		{972, 1576},
	}
	for _, c := range cases {
		v := SumDivisors(c.in)
		if v != c.out {
			t.Errorf("SumDivisors: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestNumDivisors(t *testing.T) {
	cases := []struct {
		in  int
		out int
	}{
		{1, 1},
		{2, 2},
		{3, 2},
		{4, 3},
		{5, 2},
		{6, 4},
		{7, 2},
		{8, 4},
		{9, 3},
		{10, 4},
		{11, 2},
		{12, 6},
		{13, 2},
		{14, 4},
		{15, 4},
		{16, 5},
	}
	for _, c := range cases {
		v := NumDivisors(c.in)
		if v != c.out {
			t.Errorf("NumDivisors: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestPerms(t *testing.T) {
	cases := []struct {
		in  []int
		out []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3, 1, 3, 2, 2, 1, 3, 2, 3, 1, 3, 1, 2, 3, 2, 1}},
		{[]int{2, 3, 5, 7}, []int{2, 3, 5, 7, 2, 3, 7, 5, 2, 5, 3, 7, 2, 5, 7, 3, 2,
			7, 3, 5, 2, 7, 5, 3, 3, 2, 5, 7, 3, 2, 7, 5, 3, 5, 2, 7, 3, 5, 7, 2, 3, 7,
			2, 5, 3, 7, 5, 2, 5, 2, 3, 7, 5, 2, 7, 3, 5, 3, 2, 7, 5, 3, 7, 2, 5, 7, 2,
			3, 5, 7, 3, 2, 7, 2, 3, 5, 7, 2, 5, 3, 7, 3, 2, 5, 7, 3, 5, 2, 7, 5, 2, 3,
			7, 5, 3, 2}},
	}
	for _, c := range cases {
		var i int
		for arry := range Perms(c.in) {
			for _, v := range arry {
				if v != c.out[i] {
					t.Errorf("Perms: %v\tExpected: %v", v, c.out[i])
				}
				i++
			}
		}
		if i != len(c.out) {
			t.Errorf("Perms: %v\tExpected: %v", i, len(c.out))
		}
	}
}

func TestDivmod(t *testing.T) {
	cases := []struct {
		in1, in2   int
		out1, out2 int
	}{
		{0, 1, 0, 0},
		{27, 16, 1, 11},
		{30, 3, 10, 0},
		{35, 3, 11, 2},
		{16, 6, 2, 4},
		{32, 12, 2, 8},
	}
	for _, c := range cases {
		v1, v2 := Divmod(c.in1, c.in2)
		if v1 != c.out1 || v2 != c.out2 {
			t.Errorf("Divmod: %v, %v\tExpected: %v, %v", v1, v2, c.out1, c.out2)
		}
	}
}

func TestPrimeFactors(t *testing.T) {
	cases := []struct {
		in  int
		out map[int]int
	}{
		{1, map[int]int{}},
		{2, map[int]int{2: 1}},
		{64, map[int]int{2: 6}},
		{100, map[int]int{2: 2, 5: 2}},
		{401, map[int]int{401: 1}},
		{1203, map[int]int{3: 1, 401: 1}},
		{1234, map[int]int{2: 1, 617: 1}},
	}
	for _, c := range cases {
		m := PrimeFactors(c.in)
		for k, v := range m {
			if v != c.out[k] {
				t.Errorf("PrimeFactors: %v, %v\tExpected: %v, %v", k, v, k, c.out[k])
			}
		}
		if len(m) != len(c.out) {
			t.Errorf("PrimeFactors: len(%v)\tExpected: %v", len(m), len(c.out))
		}
	}
}

func TestIsPrime(t *testing.T) {
	cases := []struct {
		in  int
		out bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{1234, false},
		{1235, false},
		{15486803, true},
		{15486873, false},
		{179426447, true},
	}
	for _, c := range cases {
		if v := IsPrime((c.in)); v != c.out {
			t.Errorf("IsPrime: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestSievePrime(t *testing.T) {
	cases := []struct {
		in  int
		out []int
	}{
		{2, []int{}},
		{10, []int{2, 3, 5, 7}},
		{100, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59,
			61, 67, 71, 73, 79, 83, 89, 97}},
	}
	for _, c := range cases {
		ps := SievePrime(c.in)
		for i, p := range ps {
			if p != c.out[i] {
				t.Errorf("SievePrime: %v\tExpected: %v", p, c.out[i])
			}
		}
		if len(ps) != len(c.out) {
			t.Errorf("SievePrime: len(%v)\tExpected: len(%v)", len(ps), len(c.out))
		}
	}
}

func TestPrimeRange(t *testing.T) {
	cases := []struct {
		in1, in2 int
		out      []int
	}{
		{2, 0, nil},
		{85, 100, []int{89, 97}},
		{700, 770, []int{701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769}},
		{100, 1000, []int{101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157,
			163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239,
			241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331,
			337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421,
			431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509,
			521, 523, 541, 547, 557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613,
			617, 619, 631, 641, 643, 647, 653, 659, 661, 673, 677, 683, 691, 701, 709,
			719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797, 809, 811, 821,
			823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911, 919,
			929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997}},
	}
	for _, c := range cases {
		p := PrimeRange(c.in1, c.in2)
		for i, v := range p {
			if v != c.out[i] {
				t.Errorf("PrimeRange: %v\tExpected: %v", v, c.out[i])
			}
		}
	}
}

func TestExp(t *testing.T) {
	cases := []struct {
		in1 int
		in2 int
		out *big.Int
	}{
		{2, 2, big.NewInt(4)},
		{2, 10, big.NewInt(1024)},
		{5, 23, big.NewInt(11920928955078125)},
	}
	for _, c := range cases {
		if v := Exp(c.in1, c.in2); v.Cmp(c.out) != 0 {
			t.Errorf("Exp: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestDigitSum(t *testing.T) {
	cases := []struct {
		in1 int
		in2 int
		out *big.Int
	}{
		{12, 1, big.NewInt(3)},
		{123, 2, big.NewInt(14)},
		{23571113, 3, big.NewInt(533)},
		{1123581321, 11, big.NewInt(8639121111)},
	}
	for _, c := range cases {
		if v := DigitSum(c.in1, c.in2); v.Cmp(c.out) != 0 {
			t.Errorf("DigitSum: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestDigitSumBig(t *testing.T) {
	cases := []struct {
		in  *big.Int
		out *big.Int
	}{
		{big.NewInt(123456789), big.NewInt(45)},
		{big.NewInt(9223372036854775807), big.NewInt(88)},
	}
	for _, c := range cases {
		if v := DigitSumBig(c.in); v.Cmp(c.out) != 0 {
			t.Errorf("DigitSumBig: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestJoinInts(t *testing.T) {
	cases := []struct {
		in  []int
		out int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 1234567},
		{[]int{0, 2, 3, 0, 5, 0, 7}, 230507},
		{[]int{2, 3, 5, 7, 1}, 23571},
	}
	for _, c := range cases {
		if v := JoinInts(c.in); v != c.out {
			t.Errorf("JoinInts: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestJoinIntsBig(t *testing.T) {
	cases := []struct {
		in  []int
		out *big.Int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, big.NewInt(1234567)},
		{[]int{0, 2, 3, 0, 5, 0, 7}, big.NewInt(230507)},
		{[]int{2, 3, 5, 7, 1}, big.NewInt(23571)},
		{[]int{2, 2, 4, 6, 3, 8, 7, 2, 3, 6, 9, 2, 8, 7, 8, 5, 7, 1}, big.NewInt(224638723692878571)},
	}
	for _, c := range cases {
		if v := JoinIntsBig(c.in); v.Cmp(c.out) != 0 {
			t.Errorf("JoinIntsBig: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestProdInts(t *testing.T) {
	cases := []struct {
		in  []int
		out int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 720},
		{[]int{0, 2, 3, 4, 5, 6, 3, 12}, 0},
		{[]int{2, 3, 5, 7, 11, 13}, 30030},
		{[]int{12, 46, 78, 9, 2, 13}, 10075104},
	}
	for _, c := range cases {
		if v := ProdInts(c.in); v != c.out {
			t.Errorf("ProdInts: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestMinInt(t *testing.T) {
	cases := []struct {
		in  []int
		out int
	}{
		{[]int{2, 3, 5, 7, 11, 13, 17, 19, 23}, 2},
		{[]int{2, 3, 4, -1}, -1},
		{[]int{1, 3, 1, 0, 2, 689}, 0},
	}
	for _, c := range cases {
		if v := MinInt(c.in...); v != c.out {
			t.Errorf("MintInt: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestMaxInt(t *testing.T) {
	cases := []struct {
		in  []int
		out int
	}{
		{[]int{2, 3, 5, 7, 11, 13, 17, 19, 23}, 23},
		{[]int{2, 3, 4, -1}, 4},
		{[]int{1, 3, 1, 0, 2, 689}, 689},
	}
	for _, c := range cases {
		if v := MaxInt(c.in...); v != c.out {
			t.Errorf("MaxtInt: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestSimplifyFraction(t *testing.T) {
	cases := []struct {
		in1, in2   int
		out1, out2 int
	}{
		{2, 4, 1, 2},
		{8, 12, 2, 3},
		{10, 35, 2, 7},
		{24, 108, 2, 9},
	}
	for _, c := range cases {
		v1, v2 := SimplifyFraction(c.in1, c.in2)
		if v1 != c.out1 || v2 != c.out2 {
			t.Errorf("SimplifyFraction: %v/%v\tExpected: %v/%v", v1, v2, c.out1, c.out2)
		}
	}
}

func TestDigits(t *testing.T) {
	cases := []struct {
		in  int
		out []int
	}{
		{0, []int{0}},
		{123, []int{1, 2, 3}},
		{2357111317, []int{2, 3, 5, 7, 1, 1, 1, 3, 1, 7}},
	}
	for _, c := range cases {
		ds := Digits(c.in)
		for i, v := range ds {
			if v != c.out[i] {
				t.Errorf("Digits: %v\tExpected: %v", v, c.out[i])
			}
		}
		if len(ds) != len(c.out) {
			t.Errorf("Digits: len(%v)\tExpected: len(%v)", len(ds), len(c.out))
		}
	}
}

func TestDigitsBig(t *testing.T) {
	cases := []struct {
		in  *big.Int
		out []int
	}{
		{big.NewInt(0), []int{0}},
		{big.NewInt(1234567890), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}},
		{big.NewInt(1123581321345589144), []int{1, 1, 2, 3, 5, 8, 1, 3, 2, 1, 3, 4,
			5, 5, 8, 9, 1, 4, 4}},
	}
	for _, c := range cases {
		ds := DigitsBig(c.in)
		for i, v := range ds {
			if v != c.out[i] {
				t.Errorf("DigitsBig: %v\tExpected: %v", v, c.out[i])
			}
		}
		if len(ds) != len(c.out) {
			t.Errorf("DigitsBig: len(%v)\tExpected: len(%v)", len(ds), len(c.out))
		}
	}
}

func TestDigitsIth(t *testing.T) {
	cases := []struct {
		in1 int
		in2 int
		out int
	}{
		{123456, -123, -1},
		{123456, 0, 1},
		{2357111317, 7, 3},
		{27141816, 8, -1},
		{3141516, 2, 4},
	}
	for _, c := range cases {
		if v := DigitsIth(c.in1, c.in2); v != c.out {
			t.Errorf("DigitsIth: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestReverseSliceInts(t *testing.T) {
	cases := []struct {
		in  []int
		out []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, []int{7, 6, 5, 4, 3, 2, 1}},
		{[]int{2, 3, 5, 7, 11, 13, 17}, []int{17, 13, 11, 7, 5, 3, 2}},
		{[]int{0, 72, 5, -12, 7, 4, 17}, []int{17, 4, 7, -12, 5, 72, 0}},
	}
	for _, c := range cases {
		rs := ReverseSliceInts(c.in)
		for i, v := range rs {
			if v != c.out[i] {
				t.Errorf("ReverseSliceInts: %v\tExpected: %v", v, c.out[i])
			}
		}
		if len(rs) != len(c.out) {
			t.Errorf("ReverseSliceInts: len(%v)\tExpected: len(%v)", len(rs), len(c.out))
		}
	}
}

func TestReverseInt(t *testing.T) {
	cases := []struct {
		in  int
		out int
	}{
		{0, 0},
		{1, 1},
		{2357111317, 7131117532},
		{112358132134, 431231853211},
	}
	for _, c := range cases {
		if v := ReverseInt(c.in); v != c.out {
			t.Errorf("ReverseInt: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestReverseIntBig(t *testing.T) {
	cases := []struct {
		in  *big.Int
		out *big.Int
	}{
		{big.NewInt(0), big.NewInt(0)},
		{big.NewInt(1234567890987654321), big.NewInt(1234567890987654321)},
		{big.NewInt(1123581321345589144), big.NewInt(4419855431231853211)},
	}
	for _, c := range cases {
		if v := ReverseIntBig(c.in); v.Cmp(c.out) != 0 {
			t.Errorf("ReverseIntBig: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestFactorial(t *testing.T) {
	cases := []struct {
		in  int
		out *big.Int
	}{
		{0, big.NewInt(1)},
		{1, big.NewInt(1)},
		{2, big.NewInt(2)},
		{10, big.NewInt(3628800)},
		{15, big.NewInt(1307674368000)},
		{20, big.NewInt(2432902008176640000)},
	}
	for _, c := range cases {
		if v := Factorial(c.in); v.Cmp(c.out) != 0 {
			t.Errorf("Factorial: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestIsPalindromeInt(t *testing.T) {
	cases := []struct {
		in  int
		out bool
	}{
		{0, true},
		{12321, true},
		{1231, false},
		{112353211, true},
	}
	for _, c := range cases {
		if v := IsPalindromeInt(c.in); v != c.out {
			t.Errorf("IsPalindromeInt: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestIsPalindromeIntBig(t *testing.T) {
	cases := []struct {
		in  *big.Int
		out bool
	}{
		{big.NewInt(0), true},
		{big.NewInt(12321), true},
		{big.NewInt(1231), false},
		{big.NewInt(112353211), true},
		{big.NewInt(314159265562951413), true},
	}
	for _, c := range cases {
		if v := IsPalindromeIntBig(c.in); v != c.out {
			t.Errorf("IsPalindromeIntBig: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestIsPandigital(t *testing.T) {
	cases := []struct {
		in  int
		out bool
	}{
		{0, false},
		{1, true},
		{312, true},
		{3142, true},
		{53142, true},
		{536142, true},
		{5376147, false},
		{537861947, false},
		{537861942, true},
	}
	for _, c := range cases {
		if v := IsPandigital(c.in); v != c.out {
			t.Errorf("IsPandigital: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestIsPermuted(t *testing.T) {
	cases := []struct {
		in1 int
		in2 int
		out bool
	}{
		{0, 0, true},
		{123, 213, true},
		{1243, 213, false},
		{11235, 53121, true},
		{112035, 53121, false},
		{112035, 503121, true},
	}
	for _, c := range cases {
		if v := IsPermuted(c.in1, c.in2); v != c.out {
			t.Errorf("IsPermuted: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestTriangleNumber(t *testing.T) {
	cases := []struct {
		in  int
		out int
	}{
		{0, 0},
		{1, 1},
		{2, 3},
		{3, 6},
		{4, 10},
		{5, 15},
		{6, 21},
		{7, 28},
		{8, 36},
		{9, 45},
		{10, 55},
		{28, 406},
	}
	for _, c := range cases {
		if v := TriangleNumber(c.in); v != c.out {
			t.Errorf("TriangleNumber: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestIsTriangleNumber(t *testing.T) {
	cases := []struct {
		in  int
		out bool
	}{
		{0, true},
		{1, true},
		{3, true},
		{46, false},
		{56, false},
		{407, false},
	}
	for _, c := range cases {
		if v := IsTriangleNumber(c.in); v != c.out {
			t.Errorf("IsTriangleNumber: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestSquareNumber(t *testing.T) {
	cases := []struct {
		in  int
		out int
	}{
		{0, 0},
		{1, 1},
		{2, 4},
		{3, 9},
		{4, 16},
		{5, 25},
		{6, 36},
		{7, 49},
		{8, 64},
		{9, 81},
		{10, 100},
	}
	for _, c := range cases {
		if v := SquareNumber(c.in); v != c.out {
			t.Errorf("SquareNumber: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestIsSquareNumber(t *testing.T) {
	cases := []struct {
		in  int
		out bool
	}{
		{0, true},
		{1, true},
		{4, true},
		{10, false},
		{17, false},
		{1468944, true},
		{1874160, false},
	}
	for _, c := range cases {
		if v := IsSquareNumber(c.in); v != c.out {
			t.Errorf("IsSquareNumber: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestPentagonNumber(t *testing.T) {
	cases := []struct {
		in  int
		out int
	}{
		{0, 0},
		{1, 1},
		{2, 5},
		{3, 12},
		{4, 22},
		{5, 35},
		{6, 51},
		{7, 70},
		{8, 92},
		{9, 117},
		{10, 145},
	}
	for _, c := range cases {
		if v := PentagonNumber(c.in); v != c.out {
			t.Errorf("PentagonNumber: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestIsPentagonNumber(t *testing.T) {
	cases := []struct {
		in  int
		out bool
	}{
		{0, false},
		{1, true},
		{5, true},
		{52, false},
		{71, false},
		{3290, true},
		{4187, true},
		{4032, false},
	}
	for _, c := range cases {
		if v := IsPentagonNumber(c.in); v != c.out {
			t.Errorf("IsPentagonNumber: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestHexagonNumber(t *testing.T) {
	cases := []struct {
		in  int
		out int
	}{
		{0, 0},
		{1, 1},
		{2, 6},
		{3, 15},
		{4, 28},
		{5, 45},
		{6, 66},
		{7, 91},
		{8, 120},
		{9, 153},
		{10, 190},
	}
	for _, c := range cases {
		if v := HexagonNumber(c.in); v != c.out {
			t.Errorf("HexagonNumber: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestIsHexagonNumber(t *testing.T) {
	cases := []struct {
		in  int
		out bool
	}{
		{0, false},
		{1, true},
		{28, true},
		{154, false},
		{189, false},
		{861, true},
		{946, true},
		{947, false},
	}
	for _, c := range cases {
		if v := IsHexagonNumber(c.in); v != c.out {
			t.Errorf("IsHexagonNumber: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestHeptagonalNumber(t *testing.T) {
	cases := []struct {
		in  int
		out int
	}{
		{0, 0},
		{1, 1},
		{2, 7},
		{3, 18},
		{4, 34},
		{5, 55},
		{6, 81},
		{7, 112},
		{8, 148},
		{9, 189},
		{10, 235},
		{27, 1782},
	}
	for _, c := range cases {
		if v := HeptagonalNumber(c.in); v != c.out {
			t.Errorf("HeptagonalNumber: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestIsHeptagonalNumber(t *testing.T) {
	cases := []struct {
		in  int
		out bool
	}{
		{0, false},
		{1, true},
		{55, true},
		{111, false},
		{149, false},
		{235, true},
		{1782, true},
		{1783, false},
	}
	for _, c := range cases {
		if v := IsHeptagonalNumber(c.in); v != c.out {
			t.Errorf("IsHeptagonalNumber: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestOctagonalNumber(t *testing.T) {
	cases := []struct {
		in  int
		out int
	}{
		{0, 0},
		{1, 1},
		{2, 8},
		{3, 21},
		{4, 40},
		{5, 65},
		{6, 96},
		{7, 133},
		{8, 176},
		{9, 225},
		{10, 280},
		{18, 936},
	}
	for _, c := range cases {
		if v := OctagonalNumber(c.in); v != c.out {
			t.Errorf("OctagonalNumber: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestIsOctagonalNumber(t *testing.T) {
	cases := []struct {
		in  int
		out bool
	}{
		{0, false},
		{1, true},
		{65, true},
		{67, false},
		{95, false},
		{176, true},
		{936, true},
		{939, false},
	}
	for _, c := range cases {
		if v := IsOctagonalNumber(c.in); v != c.out {
			t.Errorf("IsOctagonalNumber: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestIsPermute(t *testing.T) {
	cases := []struct {
		in  []int
		out bool
	}{
		{[]int{1, 1, 1}, true},
		{[]int{1, 1, 10}, false},
		{[]int{123, 321, 213}, true},
		{[]int{123, 3201, 213}, false},
		{[]int{23418715, 41572138, 57413281}, true},
		{[]int{2348715, 41572138, 57413281}, false},
	}
	for _, c := range cases {
		if v := IsPermute(c.in...); v != c.out {
			t.Errorf("IsPermute: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestNCR(t *testing.T) {
	cases := []struct {
		in1 int
		in2 int
		out int64
	}{
		{12, 5, 792},
		{37, 13, 3562467300},
		{49, 25, 63205303218876},
		{52, 16, 10363194502115},
	}
	for _, c := range cases {
		if v := NCR(c.in1, c.in2); v.Cmp(big.NewInt(c.out)) != 0 {
			t.Errorf("NCR: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestCombIndex(t *testing.T) {
	cases := []struct {
		in1 int
		in2 int
		out [][]int
	}{
		{5, 2, [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 2}, {1, 3}, {1, 4}, {2, 3},
			{2, 4}, {3, 4}}},
		{7, 3, [][]int{{0, 1, 2}, {0, 1, 3}, {0, 1, 4}, {0, 1, 5}, {0, 1, 6},
			{0, 2, 3}, {0, 2, 4}, {0, 2, 5}, {0, 2, 6}, {0, 3, 4}, {0, 3, 5}, {0, 3, 6},
			{0, 4, 5}, {0, 4, 6}, {0, 5, 6}, {1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 2, 6},
			{1, 3, 4}, {1, 3, 5}, {1, 3, 6}, {1, 4, 5}, {1, 4, 6}, {1, 5, 6}, {2, 3, 4},
			{2, 3, 5}, {2, 3, 6}, {2, 4, 5}, {2, 4, 6}, {2, 5, 6}, {3, 4, 5}, {3, 4, 6},
			{3, 5, 6}, {4, 5, 6}}},
	}
	for _, c := range cases {
		k := 0
		for indices := range CombIndex(c.in1, c.in2) {
			expect := c.out[k]
			if len(indices) != len(expect) {
				t.Errorf("CombIndex: len(%v)\tExpected: len(%v)", len(indices), len(expect))
			}
			for i, v := range indices {
				if v != expect[i] {
					t.Errorf("CombIndex: %v\tExpected: %v", v, expect[i])
				}
			}
			k++
		}
		if k != len(c.out) {
			t.Errorf("CombIndex: len(%v)\tExpected: len(%v)", k, len(c.out))
		}
	}
}

func TestComb(t *testing.T) {
	cases := []struct {
		in1 string
		in2 int
		out []string
	}{
		{"abc", 2, []string{"ab", "ac", "bc"}},
		{"jacky", 3, []string{"jac", "jak", "jay", "jck", "jcy", "jky", "ack", "acy",
			"aky", "cky"}},
	}
	for _, c := range cases {
		k := 0
		for slice := range Comb(c.in1, c.in2) {
			expect := c.out[k]
			for i, v := range slice {
				s := string(v.(uint8))
				e := string(expect[i])
				if s != e {
					t.Errorf("Comb: %v\tExpected: %v", s, e)
				}
			}
			k++
		}
		if k != len(c.out) {
			t.Errorf("Comb: len(%v)\tExpected: len(%v)", k, len(c.out))
		}
	}
}

func TestSqrtExpand(t *testing.T) {
	cases := []struct {
		in  int
		out []int
	}{
		{2, []int{1, 2}},
		{3, []int{1, 1, 2}},
		{17, []int{4, 8}},
		{64, []int{8}},
		{144, []int{12}},
		{234, []int{15, 3, 2, 1, 2, 1, 2, 3, 30}},
		{3889, []int{62, 2, 1, 3, 4, 2, 1, 7, 1, 1, 1, 1, 1, 17, 5, 7, 7, 5, 17, 1,
			1, 1, 1, 1, 7, 1, 2, 4, 3, 1, 2, 124}},
	}
	for _, c := range cases {
		ex := SqrtExapnd(c.in)
		if len(ex) != len(c.out) {
			t.Errorf("SqrtExpand: len(%v)\tExpected: len(%v)", len(ex), len(c.out))
		}
		for i, v := range ex {
			if v != c.out[i] {
				t.Errorf("SqrtExpand: %v\tExpected: %v", v, c.out[i])
			}
		}
	}
}

func TestConvergentSqrt(t *testing.T) {
	cases := []struct {
		in1, in2   int
		out1, out2 *big.Int
	}{
		{-2, 3, big.NewInt(0), big.NewInt(0)},
		{123, -1, big.NewInt(0), big.NewInt(0)},
		{123, 0, big.NewInt(11), big.NewInt(1)},
		{123, 7, big.NewInt(1772148577), big.NewInt(159789256)},
		{144, 3, big.NewInt(0), big.NewInt(0)},
		{3889, 12, big.NewInt(630603), big.NewInt(10112)},
		{1123581321, 23, big.NewInt(24321622557433), big.NewInt(725588330)},
	}
	for _, c := range cases {
		if x, y := ConvergentSqrt(c.in1, c.in2); x.Cmp(c.out1) != 0 || y.Cmp(c.out2) != 0 {
			t.Errorf("ConvergentSqrt: (%v, %v) -> (%v, %v)\tExpected: (%v, %v)", c.in1,
				c.in2, x, y, c.out1, c.out2)
		}
	}
}

func TestPellFundamental(t *testing.T) {
	cases := []struct {
		in   int
		out1 int64
		out2 int64
	}{
		{2, 3, 2},
		{3, 2, 1},
		{5, 9, 4},
		{6, 5, 2},
		{7, 8, 3},
		{8, 3, 1},
		{10, 19, 6},
		{11, 10, 3},
		{12, 7, 2},
		{13, 649, 180},
		{14, 15, 4},
		{15, 4, 1},
		{81, 1, 0},
	}
	for _, c := range cases {
		v1, v2 := PellFundamental(c.in)
		if v1.Int64() != c.out1 || v2.Int64() != c.out2 {
			t.Errorf("PellFundamental: %v %v\tExpected: %v %v", v1, v2, c.out1, c.out2)
		}
	}
}

func TestGCD(t *testing.T) {
	cases := []struct {
		in1, in2 int
		out      int
	}{
		{0, 0, -1},
		{8, 12, 4},
		{54, 24, 6},
		{18, 48, 6},
		{42, 56, 14},
		{42, -56, 14},
		{-35, 280, 35},
		{252, 105, 21},
	}
	for _, c := range cases {
		v := GCD(c.in1, c.in2)
		if v != c.out {
			t.Errorf("GCP: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestSqrt(t *testing.T) {
	cases := []struct {
		in1, in2 int
		out      string
	}{
		{2, 60, "141421356237309504880168872420969807856967187537694807317667"},
		{3, 60, "173205080756887729352744634150587236694280525381038062805580"},
		{5, 60, "223606797749978969640917366873127623544061835961152572427089"},
		{10, 60, "316227766016837933199889354443271853371955513932521682685750"},
	}
	for _, c := range cases {
		v := Sqrt(c.in1, c.in2).String()
		if v != c.out {
			t.Errorf("Sqrt: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestAbsInt(t *testing.T) {
	cases := []struct {
		in  int
		out int
	}{
		{2, 2},
		{-2, 2},
		{-3, 3},
		{12, 12},
		{-3889, 3889},
	}
	for _, c := range cases {
		v := AbsInt(c.in)
		if v != c.out {
			t.Errorf("AbsInt: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestCartProduct(t *testing.T) {
	s := 8

	var ref [][]int
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			for k := 0; k < s; k++ {
				for u := 0; u < s; u++ {
					for v := 0; v < s; v++ {
						for w := 0; w < s; w++ {
							ref = append(ref, []int{i, j, k, u, v, w})
						}
					}
				}
			}
		}
	}

	var i int
	for p := range CartProduct(s, 6) {
		for j, v := range p {
			if ref[i][j] != v {
				t.Errorf("CartProduct: %v\tExpected: %v", v, ref[i][j])
			}
		}
		i++
	}

	one := <-CartProduct(1, 1)
	if len(one) != 1 || one[0] != 0 {
		t.Errorf("CartProduct: %v\tExpected: [0]", one)
	}
}

func TestIncludesInt(t *testing.T) {
	cases := []struct {
		in1 []int
		in2 int
		out bool
	}{
		{[]int{1, 2, 3}, 2, true},
		{[]int{1, 2, 3}, 4, false},
		{[]int{1, 2, 3}, 8, false},
	}
	for _, c := range cases {
		v := IncludesInt(c.in1, c.in2)
		if v != c.out {
			t.Errorf("Includes: %v\tExpected: %v", v, c.out)
		}
	}
}

func TestSetInt(t *testing.T) {
	cases := []struct {
		in  []int
		out []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 1, 0, 5, 5}, []int{0, 1, 5}},
		{[]int{2, 0, 4, 5, 4}, []int{0, 2, 4, 5}},
	}
	for _, c := range cases {
		d := SetInt(c.in)
		sort.Ints(d)
		sort.Ints(c.out)
		match := true
		for i, v := range d {
			if v != c.out[i] {
				match = false
				break
			}
		}
		if !match {
			t.Errorf("SetInt: %v\tExpected: %v", d, c.out)
		}
	}
}

func TestLagrangePoly(t *testing.T) {
	cases := []struct {
		in1 []float64
		in2 []float64
		in3 []float64
		out []float64
	}{
		{[]float64{1, 2, 3}, []float64{1, 4, 9},
			[]float64{-1, -2, -3, 4, 5}, []float64{1, 4, 9, 16, 25}},
		{[]float64{1}, []float64{1}, []float64{1, 2}, []float64{1, 1}},
		{[]float64{1, 2}, []float64{1, 8}, []float64{1, 2, 3}, []float64{1, 8, 15}},
		{[]float64{1, 2, 3}, []float64{1, 8, 27},
			[]float64{-1, -2, -3, 1, 2, 3, 4}, []float64{23, 52, 93, 1, 8, 27, 58}},
		{[]float64{1, 2, 3, 4}, []float64{1, 8, 27, 64},
			[]float64{1, 2, 3, 4, 5}, []float64{1, 8, 27, 64, 125}},
	}
	for _, c := range cases {
		p := LagrangePoly(c.in1, c.in2)
		for i, v := range c.in3 {
			if p(v) != c.out[i] {
				t.Errorf("LagrangePoly: %v\tExpected: %v", p(v), c.out[i])
			}
		}
	}
}

func TestCatalan(t *testing.T) {
	cases := []struct {
		in  int64
		out int64
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 5},
		{4, 14},
		{5, 42},
		{6, 132},
		{7, 429},
		{8, 1430},
		{9, 4862},
		{10, 16796},
		{11, 58786},
		{12, 208012},
		{13, 742900},
		{14, 2674440},
		{15, 9694845},
		{16, 35357670},
		{17, 129644790},
		{18, 477638700},
		{19, 1767263190},
		{20, 6564120420},
		{21, 24466267020},
		{22, 91482563640},
		{23, 343059613650},
		{24, 1289904147324},
		{25, 4861946401452},
	}
	for _, c := range cases {
		v := Catalan(c.in).Int64()
		if v != c.out {
			t.Errorf("Catalan: %v\tExpected: %v", v, c.out)
		}
	}
}
