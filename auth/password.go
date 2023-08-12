package auth

import (
	"fmt"
	"net/http"

	"github.com/coffemanfp/docucentertest/server/errors"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hash of the provided password.
func HashPassword(password string) (string, error) {
	// Generate a bcrypt hash from the given password using the lowest cost (MinCost) for better performance.
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		err = fmt.Errorf("failed to generate password: %s", err)
	}
	return string(bytes), err
}

// CompareHashAndPassword compares a bcrypt hash with a plain password.
func CompareHashAndPassword(hashed, password string) (err error) {
	// Compare the provided bcrypt hash with the plain password.
	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		// If the comparison fails, create an error indicating password mismatch.
		err = fmt.Errorf("the password mismatched: %s", err)

		// Additionally, create a custom HTTP error using the 'errors' package, indicating unauthorized access.
		err = errors.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	return
}
