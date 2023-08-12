package utils

import (
	"html"
	"strings"
)

// RemoveSpaceAndConvertSpecialChars removes leading and trailing white spaces from the input string
// and escapes special characters using HTML escape sequences.
func RemoveSpaceAndConvertSpecialChars(s string) string {
	// Trim leading and trailing white spaces from the input string.
	trimmed := strings.TrimSpace(s)

	// Escape special characters in the trimmed string using HTML escape sequences.
	escaped := html.EscapeString(trimmed)

	return escaped
}
