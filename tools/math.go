package tools

import (
	"math"
	"math/big"
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

// JoinInts joins slice of ints and return it as an int.
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

// JoinIntsBig joins slice of ints into a big.Int.
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

// MinInt returns the smaller number of a and b.
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt returns the larger number of a and b.
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// SimplifyFraction simplifies a/b and return the (numerator, denominator).
func SimplifyFraction(a, b int) (int, int) {
	pa := PrimeFactors(a)
	pb := PrimeFactors(b)
	for k, v := range pa {
		v2, ok := pb[k]
		if !ok {
			continue
		}
		m := MinInt(v, v2)
		pa[k] = v - m
		pb[k] = v2 - m
	}
	sa := 1
	for k, v := range pa {
		for i := 0; i < v; i++ {
			sa *= k
		}
	}
	sb := 1
	for k, v := range pb {
		for i := 0; i < v; i++ {
			sb *= k
		}
	}
	return sa, sb
}

// Digits returns the individual digits as a slice of int.
func Digits(n int) []int {
	var d []int
	for n > 0 {
		d = append([]int{n % 10}, d...)
		n /= 10
	}
	return d
}

// DigitsBig returns the individual digits of a big.Int as a slice of int.
func DigitsBig(n *big.Int) []int {
	var d []int
	x := big.NewInt(0)
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

// ReverseInt reverses a given int.
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
	return JoinInts(da) == JoinInts(db)
}

// TriangleNumber returns the n-th triangle number.
func TriangleNumber(n int) int {
	return n * (n + 1) / 2
}

// IsTriangleNumber determines if a number is a triangle number.
func IsTriangleNumber(n int) bool {
	a := 2 * n
	b := SqrtInt(a)
	return b*(b+1) == a
}

// PentagonNumber returns the n-th pentagonal number.
func PentagonNumber(n int) int {
	return n * (3*n - 1) / 2
}

// IsPentagonNumber determines if a number is a pentagonal number.
func IsPentagonNumber(n int) bool {
	t := (math.Sqrt(24*float64(n)+1) + 1) / 6
	return t == math.Floor(t)
}

// HexagonNumber returns the n-th hexagonal number.
func HexagonNumber(n int) int {
	return n * (2*n - 1)
}

// IsHexagonNumber determines if a number is a hexagonal number.
func IsHexagonNumber(n int) bool {
	t := (math.Sqrt(8*float64(n)+1) + 1) / 4
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
