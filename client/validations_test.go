package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUsername_ValidUsername(t *testing.T) {
	username := "valid_user123"
	err := ValidateUsername(username)
	assert.NoError(t, err)
}

func TestValidateUsername_InvalidUsername(t *testing.T) {
	invalidUsernames := []string{
		"invalid username",
		"!@#$%^&*()",
		"very_long_username_that_exceeds_32_characters",
	}

	for _, username := range invalidUsernames {
		err := ValidateUsername(username)
		assert.Error(t, err)
	}
}
