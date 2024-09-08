package cases

import (
	"regexp"
	"strings"
)

var trainCaseRE = regexp.MustCompile("^[A-Z][a-z]+(_[A-Z][a-z]+)*$")

// ToTrainCase converts a string to Train-Case
func ToTrainCase(s string) string {
	if c := DefineStringCase(s); c != NormalCase {
		switch c {
		case TrainCase:
			return s
		case PascalCase:
			return FromPascalToTrainCase(s)
		case CamelCase:
			return FromCamelToTrainCase(s)
		case KebabCase:
			return FromKebabToTrainCase(s)
		case SnakeCase:
			return FromSnakeToTrainCase(s)
		case ScreamingSnakeCase:
			return FromScreamingCaseToTrainCase(s)
		}
	}

	return mapWordsAndJoin(s, "-", func(word string, _ int) string {
		return ToUpperFirst(word)
	})
}

// MatchTrainCase defines if the string matches the Train-Case
func MatchTrainCase(s string) bool {
	return trainCaseRE.MatchString(s)
}

// FromTrainToPascalCase converts a Train-Case string to PascalCase.
// Keep in mind that it skips spaces cause of these does not match Train-Case.
func FromTrainToPascalCase(s string) string {
	return rejoin(s, "-", "", func(s string, idx int) string {
		return s
	})
}

// FromTrainToCamelCase converts a Train-Case string to camelCase.
// Keep in mind that it skips spaces cause of these does not match Train-Case.
func FromTrainToCamelCase(s string) string {
	return rejoin(s, "-", "", func(s string, idx int) string {
		w := s
		if idx == 0 {
			w = ToLowerFirst(w)
		}
		return w
	})
}

// FromTrainToKebabCase converts a Train-Case string to kebab-case.
// Keep in mind that it skips spaces cause of these does not match Train-Case.
func FromTrainToKebabCase(s string) string {
	return rejoin(s, "-", "-", func(s string, _ int) string {
		return strings.ToLower(s)
	})
}

// FromTrainToSnakeCase converts a Train-Case string to snake_case.
// Keep in mind that it skips spaces cause of these does not match Train-Case.
func FromTrainToSnakeCase(s string) string {
	return rejoin(s, "-", "_", func(s string, _ int) string {
		return strings.ToLower(s)
	})
}

// FromTrainToScreamingSnakeCase converts a Train-Case string to SCREAMING_SNAKE_CASE.
// Keep in mind that it skips spaces cause of these does not match Train-Case.
func FromTrainToScreamingSnakeCase(s string) string {
	return rejoin(s, "-", "_", func(s string, _ int) string {
		return strings.ToUpper(s)
	})
}
