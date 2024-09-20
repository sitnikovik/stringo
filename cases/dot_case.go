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
			return FromSnakeToDotCase(s)
		case CamelCase:
			return FromCamelToDotCase(s)
		case KebabCase:
			return FromKebabToDotCase(s)
		case PascalCase:
			return FromPascalToDotCase(s)
		case ScreamingSnakeCase:
			return FromScreamingSnakeToDotCase(s)
		case TrainCase:
			return FromTrainToDotCase(s)
		}
	}

	words := SplitToWords(s)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return strings.Join(words, ".")
}
