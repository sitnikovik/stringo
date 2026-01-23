package stringo

// IndexOfRune returns the index of the first occurrence of the rune in the string.
//
// Returns -1 if the rune is not found.
//
//   - s: the string to search within.
//   - r: the rune to search for.
//
// Example:
//
//	IndexOfRune("hello", 'e') // returns 1
//	IndexOfRune("hello", 'a') // returns -1
//	IndexOfRune("hello\n", '\n') // returns 5
func IndexOfRune(s string, r rune) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}
	return -1
}
