package stringo

import (
	"strings"
	"unicode"
)

// ToPascalCase converts a string to PascalCase.
func ToPascalCase(s string) string {
	words := SplitToWords(s)
	for i, word := range words {
		words[i] = ToUpperFirst(word)
	}

	return strings.Join(words, "")
}

// ToCamelCase converts a string to camelCase
func ToCamelCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})

	for i, word := range words {
		w := strings.ToLower(word)
		if i > 0 {
			w = ToUpperFirst(w)
		}
		words[i] = w
	}

	return strings.Join(words, "")
}

// ToSnakeCase converts a string to snake_case
func ToSnakeCase(s string) string {
	words := SplitToWords(s)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return strings.Join(words, "_")
}

// ToKebabCase converts a string to kebab-case
func ToKebabCase(s string) string {
	words := SplitToWords(s)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return strings.Join(words, "-")
}

// ToUpperFirst converts the first character of a string to uppercase
func ToUpperFirst(s string) string {
	return caseFirstFunc(s, unicode.ToUpper)
}

// ToLowerFirst converts the first character of a string to lowercase
func ToLowerFirst(s string) string {
	return caseFirstFunc(s, unicode.ToLower)
}

// caseFirstFunc converts the first character of a string using the provided function
func caseFirstFunc(s string, f func(rune) rune) string {
	if len(s) == 0 {
		return ""
	}

	runes := []rune(s)
	runes[0] = f(runes[0])

	if len(runes) == 1 {
		return string(runes)
	}

	return string(runes[0]) + string(runes[1:])
}

// SplitToWords splits a string into words
func SplitToWords(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
}

// FromSnakeToPascalCase converts a snake_case string to PascalCase.
// Keep in mind that it skips spaces cause of these does not match snake_kase.
func FromSnakeToPascalCase(s string) string {
	if s == "" {
		return ""
	}

	words := strings.Split(s, "_")
	for i, word := range words {
		words[i] = ToUpperFirst((word))
	}

	return strings.Join(words, "")
}

// FromSnakeToCamelCase converts a snake_case string to camelCase.
// Keep in mind that it skips spaces cause of these does not match snake_case.
func FromSnakeToCamelCase(s string) string {
	if s == "" {
		return ""
	}

	words := strings.Split(s, "_")
	for i, word := range words {
		w := strings.ToLower(word)
		if i > 0 {
			w = ToUpperFirst(w)
		}
		words[i] = w
	}

	return strings.Join(words, "")
}

// FromSnakeToKebabCase converts a snake_case string to kebab-case.
// Keep in mind that it skips spaces cause of these does not match snake_case.
func FromSnakeToKebabCase(s string) string {
	if s == "" {
		return ""
	}

	words := strings.Split(s, "_")
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return strings.Join(words, "-")
}

// FromKebabToPascalCase converts a kebab-case string to PascalCase.
// Keep in mind that it skips spaces cause of these does not match kebab-case.
func FromKebabToPascalCase(s string) string {
	if s == "" {
		return ""
	}

	words := strings.Split(s, "-")
	for i, word := range words {
		words[i] = ToUpperFirst((word))
	}

	return strings.Join(words, "")
}

// FromKebabToCamelCase converts a kebab-case string to camelCase.
// Keep in mind that it skips spaces cause of these does not match kebab-case.
func FromKebabToCamelCase(s string) string {
	if s == "" {
		return ""
	}

	words := strings.Split(s, "-")
	for i, word := range words {
		w := strings.ToLower(word)
		if i > 0 {
			w = ToUpperFirst(w)
		}
		words[i] = w
	}

	return strings.Join(words, "")
}

// FromKebabToSnakeCase converts a kebab-case string to snake_case.
// Keep in mind that it skips spaces cause of these does not match kebab-case.
func FromKebabToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	words := strings.Split(s, "-")
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return strings.Join(words, "_")
}

// FromCamelToSnakeCase converts a camelCase string to snake_case.
// Keep in mind that it skips spaces cause of these does not match camelCase.
func FromCamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	sb := strings.Builder{}
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			sb.WriteString("_")
		}
		sb.WriteString(string(unicode.ToLower(r)))
	}

	return sb.String()
}

// FromCamelToSnakeCase converts a camelCase string to kebab-case.
// Keep in mind that it skips spaces cause of these does not match camelCase.
func FromCamelToKebabCase(s string) string {
	if s == "" {
		return ""
	}

	sb := strings.Builder{}
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			sb.WriteString("-")
		}
		sb.WriteString(string(unicode.ToLower(r)))
	}

	return sb.String()
}

// FromCamelToPascalCase converts a camelCase string to PascalCase.
// Keep in mind that it skips spaces cause of these does not match camelCase.
func FromCamelToPascalCase(s string) string {
	if s == "" {
		return ""
	}

	return ToUpperFirst(s)
}
