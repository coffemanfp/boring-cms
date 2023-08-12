package errors

import "fmt"

// Error represents a custom error type with a specific structure for database concerns.
type Error struct {
	Type    string // Type of the error
	prefix  string // Prefix for the error message
	content string // Content or details of the error
}

// Error returns a formatted error message.
func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.prefix, e.content)
}

// NewError creates a new Error instance with the provided parameters.
func NewError(t, prefix, content string) Error {
	return Error{
		Type:    t,
		prefix:  prefix,
		content: content,
	}
}
