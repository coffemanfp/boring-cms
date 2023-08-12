package product

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateVehiclePlate(t *testing.T) {
	validPlate := "ABC-123"
	invalidPlate := "123-ABC"

	t.Run("ValidPlate", func(t *testing.T) {
		err := ValidateVehiclePlate(&validPlate)
		assert.NoError(t, err)
	})

	t.Run("InvalidPlate", func(t *testing.T) {
		err := ValidateVehiclePlate(&invalidPlate)
		assert.Error(t, err)
		assert.EqualError(t, err, fmt.Sprintf("invalid vehicle plate: invalid vehicle plate format of %s", invalidPlate))
	})
}

func TestValidateGuideNumber(t *testing.T) {
	validGuideNumber := "ABC1234567"
	invalidGuideNumber := "ABC-123"

	t.Run("ValidGuideNumber", func(t *testing.T) {
		err := ValidateGuideNumber(&validGuideNumber)
		assert.NoError(t, err)
	})

	t.Run("InvalidGuideNumber", func(t *testing.T) {
		err := ValidateGuideNumber(&invalidGuideNumber)
		assert.Error(t, err)
		assert.EqualError(t, err, fmt.Sprintf("invalid guide number: invalid guide number format of %s", invalidGuideNumber))
	})
}

func TestValidatePort(t *testing.T) {
	validPort := 8080
	invalidPort := -1

	t.Run("ValidPort", func(t *testing.T) {
		err := ValidatePort(validPort)
		assert.NoError(t, err)
	})

	t.Run("InvalidPort", func(t *testing.T) {
		err := ValidatePort(invalidPort)
		assert.Error(t, err)
		assert.EqualError(t, err, fmt.Sprintf("invalid port: port must be a positive number %d", invalidPort))
	})
}

func TestValidateVault(t *testing.T) {
	validVault := 5
	invalidVault := -2

	t.Run("ValidVault", func(t *testing.T) {
		err := ValidateVault(validVault)
		assert.NoError(t, err)
	})

	t.Run("InvalidVault", func(t *testing.T) {
		err := ValidateVault(invalidVault)
		assert.Error(t, err)
		assert.EqualError(t, err, fmt.Sprintf("invalid vault: vault must be a positive number %d", invalidVault))
	})
}

func TestValidateCreator(t *testing.T) {
	validCreator := 1
	invalidCreator := 0

	t.Run("ValidCreator", func(t *testing.T) {
		err := validateCreator(validCreator)
		assert.NoError(t, err)
	})

	t.Run("InvalidCreator", func(t *testing.T) {
		err := validateCreator(invalidCreator)
		assert.Error(t, err)
		assert.EqualError(t, err, fmt.Sprintf("invalid creator id or not provided: %d", invalidCreator))
	})
}
