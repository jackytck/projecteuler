package tools

// ReverseString reverses a string
func ReverseString(s string) string {
	var r string
	for _, v := range s {
		r = string(v) + r
	}
	return r
}

// IsPalindromeString tells if a given string is a palindrome.
func IsPalindromeString(s string) bool {
	return s == ReverseString(s)
}

// IsAlpha determines if each letter of a string is an alphabet.
func IsAlpha(s string) bool {
	for i := range s {
		if s[i] < 'A' || s[i] > 'z' {
			return false
		} else if s[i] > 'Z' && s[i] < 'a' {
			return false
		}
	}
	return true
}

// WordValue computes the sum of the alphabetical position of individual letter.
func WordValue(s string) int {
	var v int
	for _, c := range s {
		if !IsAlpha(string(c)) {
			continue
		}
		v += int(c-'A') + 1
	}
	return v
}
