package cases

import "strings"

// ToScreamingSnakeCase converts a string to SCREAMING_SNAKE_CASE
func ToScreamingSnakeCase(s string) string {
	if c := DefineStringCase(s); c != NormalCase {
		switch c {
		case ScreamingSnakeCase:
			return s
		case PascalCase:
			return FromPascalToSnakeCase(s)
		case CamelCase:
			return FromCamelToSnakeCase(s)
		case KebabCase:
			return FromKebabToSnakeCase(s)
		}
	}

	words := SplitToWords(s)
	for i, word := range words {
		words[i] = strings.ToUpper(word)
	}

	return strings.Join(words, "_")
}

// MatchScreamingSnakeCase defines if the string matches the SCREAMING_SNAKE_CASE
func MatchScreamingSnakeCase(s string) bool {
	return screamingSnakeCaseRE.MatchString(s)
}
