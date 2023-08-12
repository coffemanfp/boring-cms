package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	validProduct := Product{
		ClientID:     1,
		GuideNumber:  newString("ABC1234567"),
		VehiclePlate: newString("ABC-123"),
		Port:         new(int),
		Vault:        new(int),
	}

	invalidClientID := Product{
		ClientID:     0,
		GuideNumber:  newString("ABC1234567"),
		VehiclePlate: newString("ABC-123"),
		Port:         new(int),
		Vault:        new(int),
	}

	invalidGuideNumber := Product{
		ClientID:     1,
		GuideNumber:  newString("ABC-123"),
		VehiclePlate: newString("ABC-123"),
		Port:         new(int),
		Vault:        new(int),
	}

	invalidVehiclePlate := Product{
		ClientID:     1,
		GuideNumber:  newString("ABC1234567"),
		VehiclePlate: newString("123-ABC"),
		Port:         new(int),
		Vault:        new(int),
	}

	invalidPort := Product{
		ClientID:     1,
		GuideNumber:  newString("ABC1234567"),
		VehiclePlate: newString("ABC-123"),
		Port:         new(int),
		Vault:        new(int),
	}
	*invalidPort.Port = -1

	invalidVault := Product{
		ClientID:     1,
		GuideNumber:  newString("ABC1234567"),
		VehiclePlate: newString("ABC-123"),
		Port:         new(int),
		Vault:        new(int),
	}
	*invalidVault.Vault = -2

	t.Run("ValidProduct", func(t *testing.T) {
		product, err := New(validProduct)
		assert.NoError(t, err)
		assert.Equal(t, validProduct, product)
	})

	t.Run("InvalidClientID", func(t *testing.T) {
		product, err := New(invalidClientID)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid creator id or not provided: 0")
		assert.Empty(t, product)
	})

	t.Run("InvalidGuideNumber", func(t *testing.T) {
		product, err := New(invalidGuideNumber)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid guide number format")
		assert.Empty(t, product)
	})

	t.Run("InvalidVehiclePlate", func(t *testing.T) {
		product, err := New(invalidVehiclePlate)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid vehicle plate format")
		assert.Empty(t, product)
	})

	t.Run("InvalidPort", func(t *testing.T) {
		product, err := New(invalidPort)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid port")
		assert.Empty(t, product)
	})

	t.Run("InvalidVault", func(t *testing.T) {
		product, err := New(invalidVault)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid vault")
		assert.Empty(t, product)
	})
}

func newString(s string) *string {
	n := &s
	return n
}
