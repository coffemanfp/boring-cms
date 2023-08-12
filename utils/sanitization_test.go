package utils

import (
	"testing"
)

func TestRemoveSpaceAndConvertSpecialChars(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "  This is a test  ",
			expected: "This is a test",
		},
		{
			input:    "Hello, <world> & 'universe'",
			expected: "Hello, &lt;world&gt; &amp; &#39;universe&#39;",
		},
		{
			input:    "  ",
			expected: "",
		},
		{
			input:    "",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := RemoveSpaceAndConvertSpecialChars(tc.input)
			if result != tc.expected {
				t.Errorf("Expected: %s, Got: %s", tc.expected, result)
			}
		})
	}
}
