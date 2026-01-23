package stringo

import (
	"strings"
)

// Splitunc splits string by condition
func SplitFunc(s string, cond func(rune, int) bool, incSep bool) []string {
	if cond == nil {
		return []string{s}
	}

	ss := make([]string, 0, len(s))

	sb := strings.Builder{}
	for i, r := range s {
		if cond(r, i) {
			ss = append(ss, sb.String())
			sb.Reset()
			if incSep {
				sb.WriteRune(r)
			}
			continue
		}
		sb.WriteRune(r)
	}
	if sb.Len() > 0 {
		ss = append(ss, sb.String())
	}

	return ss
}

// Between returns substring placed between two provided symbols.
//
//   - s: the source string to search in.
//   - left: the left delimiter to search for.
//   - right: the right delimiter to search for.
//
// Example:
//
//	Between("Hello, World!", ",", "!") -> " World"
func Between(s, left, right string) string {
	leftIdx := strings.Index(s, left)
	if leftIdx == -1 {
		return ""
	}
	contentStart := leftIdx + len(left)
	rightIdx := strings.Index(s[contentStart:], right)
	if rightIdx == -1 {
		return ""
	}
	rightAbsIdx := contentStart + rightIdx
	return s[contentStart:rightAbsIdx]
}
