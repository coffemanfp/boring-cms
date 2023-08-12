package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPError_Error(t *testing.T) {
	err := HTTPError{
		Code:    404,
		Message: "Not Found",
	}

	assert.Equal(t, err.Message, err.Error())
}

func TestNewHTTPError(t *testing.T) {
	t.Run("NoArguments", func(t *testing.T) {
		err := NewHTTPError(400, "Bad Request")
		httpErr, ok := err.(HTTPError)
		assert.True(t, ok)
		assert.Equal(t, 400, httpErr.Code)
		assert.Equal(t, "Bad Request", httpErr.Message)
	})

	t.Run("WithArguments", func(t *testing.T) {
		err := NewHTTPError(500, "Internal Server Error: %s", "something went wrong")
		httpErr, ok := err.(HTTPError)
		assert.True(t, ok)
		assert.Equal(t, 500, httpErr.Code)
		assert.Equal(t, "Internal Server Error: something went wrong", httpErr.Message)
	})
}
