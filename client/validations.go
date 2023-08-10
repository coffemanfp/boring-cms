package client

import (
	"fmt"
	"regexp"
)

var nicknameRegex = regexp.MustCompile(`^[a-z0-9_-]{3,32}$`)

// ValidateUsername validate the username with a regular expression.
//
//	@param username string: username to validate.
//	 @return err error: don't match the regex with the string provided.
func ValidateUsername(username string) (err error) {
	if !nicknameRegex.MatchString(username) {
		err = fmt.Errorf("invalid username: invalid username format of %s", username)
	}
	return
}
