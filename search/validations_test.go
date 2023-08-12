package search

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidatePriceRange(t *testing.T) {
	t.Run("ValidRange", func(t *testing.T) {
		err := validatePriceRange(10.0, 20.0)
		assert.NoError(t, err)
	})

	t.Run("InvalidRange", func(t *testing.T) {
		err := validatePriceRange(30.0, 20.0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "start price must be less than end price")
	})
}

func TestValidateQuantityRange(t *testing.T) {
	t.Run("ValidRange", func(t *testing.T) {
		err := validateQuantityRange(5, 10)
		assert.NoError(t, err)
	})

	t.Run("InvalidRange", func(t *testing.T) {
		err := validateQuantityRange(15, 10)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "start quantity must be less than end quantity")
	})
}

func TestValidateDeliveredAtRange(t *testing.T) {
	startTime := time.Now()
	endTime := startTime.Add(24 * time.Hour)

	t.Run("ValidRange", func(t *testing.T) {
		err := validateDeliveredAtRange(startTime, endTime)
		assert.NoError(t, err)
	})

	t.Run("InvalidRange", func(t *testing.T) {
		err := validateDeliveredAtRange(endTime, startTime)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "start delivered datetime must be earlier than end delivered datetime")
	})
}

func TestValidateJoinedAtRange(t *testing.T) {
	startTime := time.Now()
	endTime := startTime.Add(24 * time.Hour)

	t.Run("ValidRange", func(t *testing.T) {
		err := validateJoinedAtRange(startTime, endTime)
		assert.NoError(t, err)
	})

	t.Run("InvalidRange", func(t *testing.T) {
		err := validateJoinedAtRange(endTime, startTime)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "start joined datetime must be earlier than end joined datetime")
	})
}
