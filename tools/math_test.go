package tools

import (
	"sort"
	"testing"
)

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
