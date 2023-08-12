package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError_Error(t *testing.T) {
	err := Error{
		Type:    "database",
		prefix:  "not_found",
		content: "record not found",
	}
	expected := "not_found: record not found"
	assert.Equal(t, expected, err.Error())
}
