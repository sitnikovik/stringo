package stringo

import "strconv"

// Numeronym converts a string into its numeronym form.
//
// A numeronym is a word where the middle letters are replaced by the count of those letters.
// For example, "kubernetes" becomes "k8s" because there are 8 letters between 'k' and 's'.
// If the input string has less than 3 characters, it is returned unchanged.
//
// This function correctly handles Unicode strings.
//
// If you require better performance and are sure the input is ASCII, consider using NumeronymASCII().
func Numeronym(s string) string {
	rr := []rune(s)
	n := len(rr)
	if n < 3 {
		return s
	}
	num := []rune(strconv.Itoa(n - 2))
	res := make([]rune, len(num)+2)
	res[0] = rr[0]
	for i, r := range num {
		res[i+1] = r
	}
	res[len(res)-1] = rr[n-1]
	return string(res)
}

// NumeronymASCII converts an ASCII string into its numeronym form.
//
// A numeronym is a word where the middle letters are replaced by the count of those letters.
// For example, "kubernetes" becomes "k8s" because there are 8 letters between 'k' and 's'.
// If the input string has less than 3 characters, it is returned unchanged.
//
// This function is optimized for ASCII strings and may not handle Unicode characters correctly.
//
// If you need to handle Unicode strings, consider using Numeronym().
func NumeronymASCII(s string) string {
	n := len(s)
	if n < 3 {
		return s
	}
	num := []byte(strconv.Itoa(n - 2))
	res := make([]byte, 0, len(num)+2)
	res = append(res, s[0])
	res = append(res, num...)
	res = append(res, s[n-1])
	return string(res)
}
