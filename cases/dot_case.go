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

// FromDotToPascalCase converts a dot.case string to PascalCase.
// Keep in mind that it skips spaces cause of these does not match dot.case.
func FromDotToPascalCase(s string) string {
	return rejoin(s, ".", "", func(s string, idx int) string {
		return ToUpperFirst(strings.ToLower(s))
	})
}

// FromDotToCamelCase converts a dot.case string to camelCase.
// Keep in mind that it skips spaces cause of these does not match dot.case.
func FromDotToCamelCase(s string) string {
	return rejoin(s, ".", "", func(s string, idx int) string {
		w := strings.ToLower(s)
		if idx > 0 {
			w = ToUpperFirst(w)
		}
		return w
	})
}

// FromDotToSnakeCase converts a dot.case string to snake_case.
// Keep in mind that it skips spaces cause of these does not match dot.case.
func FromDotToSnakeCase(s string) string {
	return rejoin(s, ".", "_", func(s string, _ int) string {
		return strings.ToLower(s)
	})
}

// FromDotToScreamingSnakeCase converts a dot.case string to SCREAMING_SNAKE_CASE.
// Keep in mind that it skips spaces cause of these does not match dot.case.
func FromDotToScreamingSnakeCase(s string) string {
	return rejoin(s, ".", "_", func(s string, _ int) string {
		return strings.ToUpper(s)
	})
}

// FromDotToTrainCase converts a dot.case string to Train-Case.
// Keep in mind that it skips spaces cause of these does not match dot.case.
func FromDotToTrainCase(s string) string {
	return rejoin(s, ".", "-", func(s string, _ int) string {
		return ToUpperFirst(s)
	})
}

// FromDotToDotCase converts a dot.case string to kebab-case.
// Keep in mind that it skips spaces cause of these does not match dot.case.
func FromDotToKebabCase(s string) string {
	return rejoin(s, ".", "-", func(s string, _ int) string {
		return strings.ToLower(s)
	})
}
