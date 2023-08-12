package client

import (
	"fmt"
	"regexp"
)

// Regular expression to match a valid username format.
var nicknameRegex = regexp.MustCompile(`^[a-z0-9_-]{3,32}$`)

// ValidateUsername checks if the provided username adheres to the valid format.
func ValidateUsername(username string) (err error) {
	// Check if the username matches the defined regular expression pattern.
	if !nicknameRegex.MatchString(username) {
		// If the username format is invalid, create an error indicating the issue.
		err = fmt.Errorf("invalid username: invalid username format of %s", username)
	}
	return
}
