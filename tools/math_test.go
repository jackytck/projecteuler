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
