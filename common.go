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
