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
