package tools

import (
	"math"
	"math/big"
	"reflect"
	"sort"
	"strconv"
)

// Sum returns the sum of its argument.
func Sum(a ...int) int {
	var s int
	for _, v := range a {
		s += v
	}
	return s
}

// SqrtInt computes the square root of n and casts to an int.
func SqrtInt(n int) int {
	return int(math.Sqrt(float64(n)))
}

// Divisors returns all the proper divisors of n.
func Divisors(n int) []int {
	d := []int{1}
	for i := 2; i <= SqrtInt(n); i++ {
		if n%i == 0 {
			d = append(d, i)
			if other := n / i; other != i {
				d = append(d, other)
			}
		}
	}
	return d
}

// SumDivisors sums all the proper divisors of n.
func SumDivisors(n int) int {
	return Sum(Divisors(n)...)
}

// NumDivisors gives the number of divisors of n.
func NumDivisors(n int) int {
	c := 1
	p := PrimeFactors(n)
	for _, v := range p {
		c *= v + 1
	}
	return c
}

// Perms return a channel of each permutation of a slice.
func Perms(slice []int) chan []int {
	c := make(chan []int)
	go func() {
		if len(slice) == 1 {
			c <- slice
		} else {
			for i, v := range slice {
				b := append([]int(nil), slice[:i]...)
				b = append(b, slice[i+1:]...)
				for p := range Perms(b) {
					c <- append([]int{v}, p...)
				}
			}
		}
		close(c)
	}()
	return c
}

// Divmod gives the quotient and remainder of integer division.
func Divmod(a, b int) (int, int) {
	return a / b, a % b
}

// PrimeFactors gives the prime factors of an int.
func PrimeFactors(n int) map[int]int {
	factors := make(map[int]int)
	d := 2
	for n > 1 {
		for n%d == 0 {
			n /= d
			factors[d]++
		}
		d++
	}
	return factors
}

// IsPrime determines if a given number is prime.
// https://en.wikipedia.org/wiki/Primality_test
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	b := SqrtInt(n)
	for i := 5; i <= b; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// SievePrime gives all the primes below n via the method of Eratosthenes.
func SievePrime(n int) []int {
	var primes []int
	comp := make([]bool, n+2)
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		for j := i * i; j < n; j += i {
			comp[j] = true
		}
	}
	for i, v := range comp {
		if i > 1 && !v && i < n {
			primes = append(primes, i)
		}
	}
	return primes
}

// PrimeRange gives the primes in range [a, b).
func PrimeRange(a, b int) []int {
	var primes []int
	if b <= a {
		return primes
	}
	ps := SievePrime(SqrtInt(b) + 1)
	size := b - a + 1
	comp := make([]bool, size)
	for _, p := range ps {
		for j := a/p*p - a; j < size; j += p {
			if j < 0 {
				continue
			}
			comp[j] = true
		}
	}
	for i, v := range comp {
		if !v && i < size {
			primes = append(primes, a+i)
		}
	}
	return primes
}

// Exp computes x to the power of y.
func Exp(x, y int) *big.Int {
	a := big.NewInt(1)
	b := big.NewInt(int64(x))
	for i := 0; i < y; i++ {
		a.Mul(a, b)
	}
	return a
}

// DigitSum sums the digit^p of a given number.
func DigitSum(n, p int) *big.Int {
	s := big.NewInt(0)
	for _, c := range strconv.Itoa(n) {
		i, _ := strconv.Atoi(string(c))
		s.Add(s, Exp(i, p))
	}
	return s
}

// DigitSumBig sums the digits of a big number.
func DigitSumBig(n *big.Int) *big.Int {
	s := big.NewInt(0)
	for _, d := range DigitsBig(n) {
		z := big.NewInt(int64(d))
		s.Add(s, z)
	}
	return s
}

// JoinInts joins slice of single digit ints and return it as an int.
func JoinInts(slice []int) int {
	var sum int
	len := len(slice) - 1
	for i, v := range slice {
		p := 1
		for j := 0; j < len-i; j++ {
			p *= 10
		}
		sum += p * v
	}
	return sum
}

// JoinIntsBig joins slice of single digit ints into a big.Int.
func JoinIntsBig(slice []int) *big.Int {
	sum := big.NewInt(0)
	ten := big.NewInt(10)
	p := Exp(10, len(slice)-1)
	for _, v := range slice {
		z := big.NewInt(1)
		z.Mul(p, big.NewInt(int64(v)))
		sum.Add(sum, z)
		p.Div(p, ten)
	}
	return sum
}

// ProdInts computes the product of the slice of ints.
func ProdInts(slice []int) int {
	p := 1
	for _, v := range slice {
		p *= v
	}
	return p
}

// MinInt returns the smallest number.
func MinInt(a ...int) int {
	var ret int
	for i, v := range a {
		if i == 0 {
			ret = v
		}
		if v < ret {
			ret = v
		}
	}
	return ret
}

// MaxInt returns the largest number.
func MaxInt(a ...int) int {
	var ret int
	for i, v := range a {
		if i == 0 {
			ret = v
		}
		if v > ret {
			ret = v
		}
	}
	return ret
}

// SimplifyFraction simplifies a/b and return the (numerator, denominator).
func SimplifyFraction(a, b int) (int, int) {
	gcd := GCD(a, b)
	return a / gcd, b / gcd
}

// Digits returns the individual digits as a slice of int.
func Digits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	var d []int
	for n > 0 {
		d = append([]int{n % 10}, d...)
		n /= 10
	}
	return d
}

// DigitsBig returns the individual digits of a big.Int as a slice of int.
func DigitsBig(n *big.Int) []int {
	x := big.NewInt(0)
	if n.Cmp(x) == 0 {
		return []int{0}
	}
	var d []int
	x.Set(n)
	zero := big.NewInt(0)
	ten := big.NewInt(10)
	for x.Cmp(zero) > 0 {
		m := big.NewInt(1)
		x.DivMod(x, ten, m)
		d = append([]int{int(m.Int64())}, d...)
	}
	return d
}

// DigitsIth returns the i-th digit of a number n.
func DigitsIth(n, i int) int {
	d := Digits(n)
	if i < 0 || i >= len(d) {
		return -1
	}
	return Digits(n)[i]
}

// ReverseSliceInts reverses a slice of ints, and return the new slice.
func ReverseSliceInts(a []int) []int {
	s := len(a)
	r := make([]int, s)
	for i, v := range a {
		r[s-1-i] = v
	}
	return r
}

// ReverseInt reverses a non-negative int.
func ReverseInt(n int) int {
	return JoinInts(ReverseSliceInts(Digits(n)))
}

// ReverseIntBig reverses a given big.Int.
func ReverseIntBig(n *big.Int) *big.Int {
	return JoinIntsBig(ReverseSliceInts(DigitsBig(n)))
}

// Factorial returns the fractorial of a number.
func Factorial(n int) *big.Int {
	a := big.NewInt(1)
	for i := 2; i <= n; i++ {
		a.Mul(a, big.NewInt(int64(i)))
	}
	return a
}

// IsPalindromeInt tells if a number is a palindrome.
func IsPalindromeInt(n int) bool {
	return IsPalindromeString(strconv.Itoa(n))
}

// IsPalindromeIntBig tells if a big.Int number is a palindrome.
func IsPalindromeIntBig(n *big.Int) bool {
	return IsPalindromeString(n.String())
}

// IsPandigital determines if a number is pandigital.
func IsPandigital(n int) bool {
	d := Digits(n)
	sort.Ints(d)
	ans := true
	for i, v := range d {
		if i+1 != v {
			ans = false
			break
		}
	}
	return ans
}

// IsPermuted determines if digits in the two given ints are permuted.
func IsPermuted(a, b int) bool {
	da := Digits(a)
	db := Digits(b)
	sort.Ints(da)
	sort.Ints(db)
	return len(da) == len(db) && JoinInts(da) == JoinInts(db)
}

// TriangleNumber returns the n-th triangle number.
// e.g. 1, 3, 6, 10, 15, ...
func TriangleNumber(n int) int {
	return n * (n + 1) / 2
}

// IsTriangleNumber determines if a number is a triangle number.
func IsTriangleNumber(n int) bool {
	return IsSquareNumber(8*n + 1)
}

// SquareNumber returns the n-th square number.
// e.g. 1, 4, 9, 16, 25, ...
func SquareNumber(n int) int {
	return n * n
}

// IsSquareNumber determines if a number is a square number.
func IsSquareNumber(n int) bool {
	t := math.Sqrt(float64(n))
	return t == math.Floor(t)
}

// PentagonNumber returns the n-th pentagonal number.
// e.g. 1, 5, 12, 22, 35, ...
func PentagonNumber(n int) int {
	return n * (3*n - 1) / 2
}

// IsPentagonNumber determines if a number is a pentagonal number.
func IsPentagonNumber(n int) bool {
	t := (math.Sqrt(24*float64(n)+1) + 1) / 6
	return t == math.Floor(t)
}

// HexagonNumber returns the n-th hexagonal number.
// e.g. 1, 6, 15, 28, 45, ...
func HexagonNumber(n int) int {
	return n * (2*n - 1)
}

// IsHexagonNumber determines if a number is a hexagonal number.
func IsHexagonNumber(n int) bool {
	t := (math.Sqrt(8*float64(n)+1) + 1) / 4
	return t == math.Floor(t)
}

// HeptagonalNumber returns the n-th heptagonal number.
// e.g. 1, 7, 18, 34, 55, ...
func HeptagonalNumber(n int) int {
	return n * (5*n - 3) / 2
}

// IsHeptagonalNumber determines if a number is a heptagonal number.
func IsHeptagonalNumber(n int) bool {
	t := (math.Sqrt(40*float64(n)+9) + 3) / 10
	return t == math.Floor(t)
}

// OctagonalNumber returns the n-th octagonal number.
// e.g. 1, 8, 21, 40, 65, ...
func OctagonalNumber(n int) int {
	return n * (3*n - 2)
}

// IsOctagonalNumber determines if a number is a octagonal number.
func IsOctagonalNumber(n int) bool {
	t := (math.Sqrt(3*float64(n)+1) + 1) / 3
	return t == math.Floor(t)
}

// IsPermute determines if all the given ints are permutation of each other.
func IsPermute(args ...int) bool {
	permute := true
	var compare int
	for i, v := range args {
		d := Digits(v)
		sort.Ints(d)
		j := JoinInts(d)
		if i == 0 {
			compare = j
		} else {
			if j != compare {
				permute = false
				break
			}
		}
	}
	return permute
}

// NCR computes the number nCr.
func NCR(n, r int) *big.Int {
	z := big.NewInt(0)
	return z.Binomial(int64(n), int64(r))
}

// CombIndex gives the combinations of selecting [0, k-1] from [0, n-1] in
// ascending order.
func CombIndex(n, k int) <-chan []int {
	c := make(chan []int)
	var comb []int
	for i := 0; i < k; i++ {
		comb = append(comb, i)
	}
	next := func() bool {
		i := k - 1
		comb[i]++
		for i > 0 && comb[i] >= n-k+1+i {
			i--
			comb[i]++
		}
		if comb[0] > n-k {
			return false
		}
		for i = i + 1; i < k; i++ {
			comb[i] = comb[i-1] + 1
		}
		return true
	}
	copy := func(a []int) []int {
		var b []int
		for _, v := range a {
			b = append(b, v)
		}
		return b
	}
	go func() {
		defer close(c)
		c <- copy(comb)
		for next() {
			c <- copy(comb)
		}
	}()
	return c
}

// Comb gives the combinations of selecting k items from the slice of collection.
func Comb(collection interface{}, k int) <-chan []interface{} {
	col := reflect.ValueOf(collection)
	c := make(chan []interface{})
	go func() {
		defer close(c)
		for indices := range CombIndex(col.Len(), k) {
			var chosen []interface{}
			for _, i := range indices {
				chosen = append(chosen, col.Index(i).Interface())
			}
			c <- chosen
		}
	}()
	return c
}

// SqrtExapnd gives the expansion of continued fraction of square root of n.
// https://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Continued_fraction_expansion
func SqrtExapnd(n int) []int {
	srt := math.Sqrt(float64(n))
	if IsSquareNumber(n) {
		return []int{int(srt)}
	}
	var terms []int
	var m int
	d := 1
	a := int(math.Floor(srt))
	a0 := a
	terms = append(terms, a)
	for a != 2*a0 {
		m = d*a - m
		d = (n - m*m) / d
		a = int(math.Floor(float64((a0 + m)) / float64(d)))
		terms = append(terms, a)
	}
	return terms
}

// ConvergentSqrt gives the i-th convergent of the continued fraction of the
// squart root of n.
func ConvergentSqrt(n, i int) (*big.Int, *big.Int) {
	if i < 0 || IsSquareNumber(n) {
		return big.NewInt(0), big.NewInt(0)
	}
	a := SqrtExapnd(n)
	if i == 0 {
		return big.NewInt(int64(a[0])), big.NewInt(1)
	}

	s := len(a) - 1
	p, q := big.NewInt(1), big.NewInt(0)
	p2, q2 := big.NewInt(0), big.NewInt(1)
	for j := 0; j <= i; j++ {
		an := a[(j-1)%s+1]
		if j == 0 {
			an = a[0]
		}

		// p2, p = p, an*p+p2
		nextP := big.NewInt(0).Set(p)
		nextP.Mul(nextP, big.NewInt(int64(an)))
		nextP.Add(nextP, p2)
		p2, p = p, nextP

		// q2, q = q, an*q+q2
		nextQ := big.NewInt(0).Set(q)
		nextQ.Mul(nextQ, big.NewInt(int64(an)))
		nextQ.Add(nextQ, q2)
		q2, q = q, nextQ
	}
	return p, q
}

// PellFundamental gives the fundamental solution (mx, my) of Pell's equaltion:
// x^2 - n * y^2 = 1, where n is a given positive nonsquare integer and integer
// solutions are sought for x and y.
func PellFundamental(n int) (*big.Int, *big.Int) {
	mx, my := big.NewInt(1), big.NewInt(0)
	if IsSquareNumber(n) {
		return mx, my
	}
	for i := 1; true; i++ {
		x, y := ConvergentSqrt(n, i)
		mx.Set(x)
		my.Set(y)
		x.Mul(x, x)
		y.Mul(y, y)
		y.Mul(y, big.NewInt(int64(-n)))
		if x.Add(x, y); x.Int64() == 1 {
			break
		}
	}
	return mx, my
}

// GCD computes the gcd of x and y.
func GCD(x, y int) int {
	if x == 0 && y == 0 {
		return -1
	}
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// Sqrt computes the square root of n with precision number of digits.
// e.g. Sqrt(2, 20) => 14142135623730950488
func Sqrt(n, precision int) *big.Int {
	limit := Exp(10, precision+1)
	a := big.NewInt(5 * int64(n))
	b := big.NewInt(5)
	ten := big.NewInt(10)
	hundred := big.NewInt(100)
	fortyFive := big.NewInt(45)
	for b.Cmp(limit) < 0 {
		if a.Cmp(b) >= 0 {
			a.Sub(a, b)
			b.Add(b, ten)
		} else {
			a.Mul(a, hundred)
			b.Mul(b, ten)
			b.Sub(b, fortyFive)
		}
	}
	return b.Div(b, hundred)
}

// AbsInt returns the absolute value of x.
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// CartProduct produces the indices of the Cartesian Product.
func CartProduct(size, repeat int) chan []int {
	c := make(chan []int)
	go func() {
		defer close(c)
		if size <= 1 || repeat < 1 {
			c <- []int{0}
		} else {
			list := make([]int, repeat)
			for {
				copy := append([]int{}, list...)
				c <- copy

				// add one to the last index
				list[repeat-1]++

				// carry
				for i := repeat - 1; i > 0; i-- {
					if list[i] == size {
						list[i] = 0
						list[i-1]++
					}
				}

				// exhausted all
				if list[0] == size {
					return
				}
			}
		}
	}()
	return c
}

// IncludesInt determines whether a slice includes a certain element.
func IncludesInt(s []int, e int) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// SetInt returns the set of int from the given slice of int.
func SetInt(s []int) []int {
	m := make(map[int]bool)
	for _, v := range s {
		m[v] = true
	}
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// LagrangePoly gives the Lagrange interpolating polynomial P(x) of
// degree <=(n-1) that passes through the n given points.
func LagrangePoly(xs []float64, ys []float64) func(float64) float64 {
	p := func(x float64) float64 {
		var s float64
		for i, u := range ys {
			m := u
			xi := xs[i]
			for j, v := range xs {
				if i == j {
					continue
				}
				m *= (x - v) / (xi - v)
			}
			s += m
		}
		return s
	}
	return p
}

// Catalan returns the Catalan Number of n.
func Catalan(n int64) *big.Int {
	z := big.NewInt(0)
	z.Binomial(2*n, n)
	return z.Quo(z, big.NewInt(n+1))
}
