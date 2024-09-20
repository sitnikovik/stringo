package cases

import (
	"regexp"
	"strings"
)

var dotCaseRE = regexp.MustCompile(`^([a-z]+\.)*[a-z]+$`)

// MatchDotCase defines if the string matches the dot.case
func MatchDotCase(s string) bool {
	return dotCaseRE.MatchString(s)
}

// ToDotCase converts a string to dot.case
func ToDotCase(s string) string {
	if c := DefineStringCase(s); c != NormalCase {
		switch c {
		case DotCase:
			return s
		case SnakeCase:
			panic("implement me")
		case CamelCase:
			panic("implement me")
		case KebabCase:
			panic("implement me")
		case PascalCase:
			panic("implement me")
		case ScreamingSnakeCase:
			panic("implement me")
		case TrainCase:
			panic("implement me")
		}
	}

	words := SplitToWords(s)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return strings.Join(words, ".")
}
