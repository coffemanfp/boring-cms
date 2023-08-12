package search

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewSearch(t *testing.T) {
	validSearch := Search{
		ClientID:      1,
		Port:          80,
		Vault:         2,
		GuideNumber:   "ABC1234567",
		VehiclePlate:  "ABC-123",
		PriceRange:    RangeFloat64{Start: 100.0, End: 200.0},
		QuantityRange: RangeInt{Start: 5, End: 10},
		JoinedAtRange: RangeTime{
			Start: parseTimeIgnoringError(time.Now().Add(-24 * time.Hour).Format(time.RFC3339)),
			End:   parseTimeIgnoringError(time.Now().Format(time.RFC3339)),
		},
		DeliveredAtRange: RangeTime{
			Start: parseTimeIgnoringError(time.Now().Add(-48 * time.Hour).Format(time.RFC3339)),
			End:   parseTimeIgnoringError(time.Now().Add(-24 * time.Hour).Format(time.RFC3339)),
		},
	}

	invalidPort := Search{
		ClientID:      1,
		Port:          -1,
		Vault:         2,
		GuideNumber:   "ABC1234567",
		VehiclePlate:  "ABC-123",
		PriceRange:    RangeFloat64{Start: 100.0, End: 200.0},
		QuantityRange: RangeInt{Start: 5, End: 10},
		JoinedAtRange: RangeTime{
			Start: parseTimeIgnoringError(time.Now().Add(-24 * time.Hour).Format(time.RFC3339)),
			End:   parseTimeIgnoringError(time.Now().Format(time.RFC3339)),
		},
		DeliveredAtRange: RangeTime{
			Start: parseTimeIgnoringError(time.Now().Add(-48 * time.Hour).Format(time.RFC3339)),
			End:   parseTimeIgnoringError(time.Now().Add(-24 * time.Hour).Format(time.RFC3339)),
		},
	}

	invalidVault := Search{
		ClientID:      1,
		Port:          80,
		Vault:         -2,
		GuideNumber:   "ABC1234567",
		VehiclePlate:  "ABC-123",
		PriceRange:    RangeFloat64{Start: 100.0, End: 200.0},
		QuantityRange: RangeInt{Start: 5, End: 10},
		JoinedAtRange: RangeTime{
			Start: parseTimeIgnoringError(time.Now().Add(-24 * time.Hour).Format(time.RFC3339)),
			End:   parseTimeIgnoringError(time.Now().Format(time.RFC3339)),
		},
		DeliveredAtRange: RangeTime{
			Start: parseTimeIgnoringError(time.Now().Add(-48 * time.Hour).Format(time.RFC3339)),
			End:   parseTimeIgnoringError(time.Now().Add(-24 * time.Hour).Format(time.RFC3339)),
		},
	}

	invalidGuideNumber := Search{
		ClientID:      1,
		Port:          80,
		Vault:         2,
		GuideNumber:   "ABC-123",
		VehiclePlate:  "ABC-123",
		PriceRange:    RangeFloat64{Start: 100.0, End: 200.0},
		QuantityRange: RangeInt{Start: 5, End: 10},
		JoinedAtRange: RangeTime{
			Start: parseTimeIgnoringError(time.Now().Add(-24 * time.Hour).Format(time.RFC3339)),
			End:   parseTimeIgnoringError(time.Now().Format(time.RFC3339)),
		},
		DeliveredAtRange: RangeTime{
			Start: parseTimeIgnoringError(time.Now().Add(-48 * time.Hour).Format(time.RFC3339)),
			End:   parseTimeIgnoringError(time.Now().Add(-24 * time.Hour).Format(time.RFC3339)),
		},
	}

	invalidVehiclePlate := Search{
		ClientID:      1,
		Port:          80,
		Vault:         2,
		GuideNumber:   "ABC1234567",
		VehiclePlate:  "123-ABC",
		PriceRange:    RangeFloat64{Start: 100.0, End: 200.0},
		QuantityRange: RangeInt{Start: 5, End: 10},
		JoinedAtRange: RangeTime{
			Start: parseTimeIgnoringError(time.Now().Add(-24 * time.Hour).Format(time.RFC3339)),
			End:   parseTimeIgnoringError(time.Now().Format(time.RFC3339)),
		},
		DeliveredAtRange: RangeTime{
			Start: parseTimeIgnoringError(time.Now().Add(-48 * time.Hour).Format(time.RFC3339)),
			End:   parseTimeIgnoringError(time.Now().Add(-24 * time.Hour).Format(time.RFC3339)),
		},
	}

	invalidPriceRange := Search{
		ClientID:      1,
		Port:          80,
		Vault:         2,
		GuideNumber:   "ABC1234567",
		VehiclePlate:  "ABC-123",
		PriceRange:    RangeFloat64{Start: 200.0, End: 100.0}, // Invalid range
		QuantityRange: RangeInt{Start: 5, End: 10},
		JoinedAtRange: RangeTime{
			Start: parseTimeIgnoringError(time.Now().Add(-24 * time.Hour).Format(time.RFC3339)),
			End:   parseTimeIgnoringError(time.Now().Format(time.RFC3339)),
		},
		DeliveredAtRange: RangeTime{
			Start: parseTimeIgnoringError(time.Now().Add(-48 * time.Hour).Format(time.RFC3339)),
			End:   parseTimeIgnoringError(time.Now().Add(-24 * time.Hour).Format(time.RFC3339)),
		},
	}

	invalidQuantityRange := Search{
		ClientID:      1,
		Port:          80,
		Vault:         2,
		GuideNumber:   "ABC1234567",
		VehiclePlate:  "ABC-123",
		PriceRange:    RangeFloat64{Start: 100.0, End: 200.0},
		QuantityRange: RangeInt{Start: 10, End: 5}, // Invalid range
		JoinedAtRange: RangeTime{
			Start: parseTimeIgnoringError(time.Now().Add(-24 * time.Hour).Format(time.RFC3339)),
			End:   parseTimeIgnoringError(time.Now().Format(time.RFC3339)),
		},
		DeliveredAtRange: RangeTime{
			Start: parseTimeIgnoringError(time.Now().Add(-48 * time.Hour).Format(time.RFC3339)),
			End:   parseTimeIgnoringError(time.Now().Add(-24 * time.Hour).Format(time.RFC3339)),
		},
	}

	invalidJoinedAtRange := Search{
		ClientID:      1,
		Port:          80,
		Vault:         2,
		GuideNumber:   "ABC1234567",
		VehiclePlate:  "ABC-123",
		PriceRange:    RangeFloat64{Start: 100.0, End: 200.0},
		QuantityRange: RangeInt{Start: 5, End: 10},
		JoinedAtRange: RangeTime{Start: time.Now(), End: time.Now().Add(-24 * time.Hour)}, // Invalid range
		DeliveredAtRange: RangeTime{
			Start: parseTimeIgnoringError(time.Now().Add(-48 * time.Hour).Format(time.RFC3339)),
			End:   parseTimeIgnoringError(time.Now().Add(-24 * time.Hour).Format(time.RFC3339)),
		},
	}

	invalidDeliveredAtRange := Search{
		ClientID:         1,
		Port:             80,
		Vault:            2,
		GuideNumber:      "ABC1234567",
		VehiclePlate:     "ABC-123",
		PriceRange:       RangeFloat64{Start: 100.0, End: 200.0},
		QuantityRange:    RangeInt{Start: 5, End: 10},
		JoinedAtRange:    RangeTime{Start: time.Now().Add(-24 * time.Hour), End: time.Now()},
		DeliveredAtRange: RangeTime{Start: time.Now().Add(-24 * time.Hour), End: time.Now().Add(-48 * time.Hour)}, // Invalid range
	}

	t.Run("ValidSearch", func(t *testing.T) {
		search, err := New(validSearch.ClientID, validSearch.Port, validSearch.Vault, validSearch.GuideNumber,
			validSearch.Type, validSearch.VehiclePlate, validSearch.PriceRange.Start, validSearch.PriceRange.End,
			validSearch.QuantityRange.Start, validSearch.QuantityRange.End, validSearch.JoinedAtRange.Start.Format(time.RFC3339),
			validSearch.JoinedAtRange.End.Format(time.RFC3339), validSearch.DeliveredAtRange.Start.Format(time.RFC3339), validSearch.DeliveredAtRange.End.Format(time.RFC3339))
		assert.NoError(t, err)
		if assert.WithinDuration(t, validSearch.JoinedAtRange.Start, search.JoinedAtRange.Start, time.Second) {
			validSearch.JoinedAtRange.Start = search.JoinedAtRange.Start
		}
		if assert.WithinDuration(t, validSearch.JoinedAtRange.End, search.JoinedAtRange.End, time.Second) {
			validSearch.JoinedAtRange.End = search.JoinedAtRange.End
		}
		if assert.WithinDuration(t, validSearch.DeliveredAtRange.Start, search.DeliveredAtRange.Start, time.Second) {
			validSearch.DeliveredAtRange.Start = search.DeliveredAtRange.Start
		}
		if assert.WithinDuration(t, validSearch.DeliveredAtRange.End, search.DeliveredAtRange.End, time.Second) {
			validSearch.DeliveredAtRange.End = search.DeliveredAtRange.End
		}
		assert.Equal(t, validSearch, search)
	})

	t.Run("InvalidPort", func(t *testing.T) {
		search, err := New(invalidPort.ClientID, invalidPort.Port, invalidPort.Vault, invalidPort.GuideNumber,
			invalidPort.Type, invalidPort.VehiclePlate, invalidPort.PriceRange.Start, invalidPort.PriceRange.End,
			invalidPort.QuantityRange.Start, invalidPort.QuantityRange.End, invalidPort.JoinedAtRange.Start.Format(time.RFC3339),
			invalidPort.JoinedAtRange.End.Format(time.RFC3339), invalidPort.DeliveredAtRange.Start.Format(time.RFC3339), invalidPort.DeliveredAtRange.End.Format(time.RFC3339))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid port")
		assert.Empty(t, search)
	})

	t.Run("InvalidVault", func(t *testing.T) {
		search, err := New(invalidVault.ClientID, invalidVault.Port, invalidVault.Vault, invalidVault.GuideNumber,
			invalidVault.Type, invalidVault.VehiclePlate, invalidVault.PriceRange.Start, invalidVault.PriceRange.End,
			invalidVault.QuantityRange.Start, invalidVault.QuantityRange.End, invalidVault.JoinedAtRange.Start.Format(time.RFC3339),
			invalidVault.JoinedAtRange.End.Format(time.RFC3339), invalidVault.DeliveredAtRange.Start.Format(time.RFC3339), invalidVault.DeliveredAtRange.End.Format(time.RFC3339))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid vault")
		assert.Empty(t, search)
	})

	t.Run("InvalidGuideNumber", func(t *testing.T) {
		search, err := New(invalidGuideNumber.ClientID, invalidGuideNumber.Port, invalidGuideNumber.Vault, invalidGuideNumber.GuideNumber,
			invalidGuideNumber.Type, invalidGuideNumber.VehiclePlate, invalidGuideNumber.PriceRange.Start, invalidGuideNumber.PriceRange.End,
			invalidGuideNumber.QuantityRange.Start, invalidGuideNumber.QuantityRange.End, invalidGuideNumber.JoinedAtRange.Start.Format(time.RFC3339),
			invalidGuideNumber.JoinedAtRange.End.Format(time.RFC3339), invalidGuideNumber.DeliveredAtRange.Start.Format(time.RFC3339), invalidGuideNumber.DeliveredAtRange.End.Format(time.RFC3339))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid guide number format")
		assert.Empty(t, search)
	})

	t.Run("InvalidVehiclePlate", func(t *testing.T) {
		search, err := New(invalidVehiclePlate.ClientID, invalidVehiclePlate.Port, invalidVehiclePlate.Vault, invalidVehiclePlate.GuideNumber,
			invalidVehiclePlate.Type, invalidVehiclePlate.VehiclePlate, invalidVehiclePlate.PriceRange.Start, invalidVehiclePlate.PriceRange.End,
			invalidVehiclePlate.QuantityRange.Start, invalidVehiclePlate.QuantityRange.End, invalidVehiclePlate.JoinedAtRange.Start.Format(time.RFC3339),
			invalidVehiclePlate.JoinedAtRange.End.Format(time.RFC3339), invalidVehiclePlate.DeliveredAtRange.Start.Format(time.RFC3339), invalidVehiclePlate.DeliveredAtRange.End.Format(time.RFC3339))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid vehicle plate format")
		assert.Empty(t, search)
	})

	t.Run("InvalidPriceRange", func(t *testing.T) {
		search, err := New(invalidPriceRange.ClientID, invalidPriceRange.Port, invalidPriceRange.Vault, invalidPriceRange.GuideNumber,
			invalidPriceRange.Type, invalidPriceRange.VehiclePlate, invalidPriceRange.PriceRange.Start, invalidPriceRange.PriceRange.End,
			invalidPriceRange.QuantityRange.Start, invalidPriceRange.QuantityRange.End, invalidPriceRange.JoinedAtRange.Start.Format(time.RFC3339),
			invalidPriceRange.JoinedAtRange.End.Format(time.RFC3339), invalidPriceRange.DeliveredAtRange.Start.Format(time.RFC3339), invalidPriceRange.DeliveredAtRange.End.Format(time.RFC3339))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid price range")
		assert.Empty(t, search)
	})

	t.Run("InvalidQuantityRange", func(t *testing.T) {
		search, err := New(invalidQuantityRange.ClientID, invalidQuantityRange.Port, invalidQuantityRange.Vault, invalidQuantityRange.GuideNumber,
			invalidQuantityRange.Type, invalidQuantityRange.VehiclePlate, invalidQuantityRange.PriceRange.Start, invalidQuantityRange.PriceRange.End,
			invalidQuantityRange.QuantityRange.Start, invalidQuantityRange.QuantityRange.End, invalidQuantityRange.JoinedAtRange.Start.Format(time.RFC3339),
			invalidQuantityRange.JoinedAtRange.End.Format(time.RFC3339), invalidQuantityRange.DeliveredAtRange.Start.Format(time.RFC3339), invalidQuantityRange.DeliveredAtRange.End.Format(time.RFC3339))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid quantity range")
		assert.Empty(t, search)
	})

	t.Run("InvalidJoinedAtRange", func(t *testing.T) {
		search, err := New(invalidJoinedAtRange.ClientID, invalidJoinedAtRange.Port, invalidJoinedAtRange.Vault, invalidJoinedAtRange.GuideNumber,
			invalidJoinedAtRange.Type, invalidJoinedAtRange.VehiclePlate, invalidJoinedAtRange.PriceRange.Start, invalidJoinedAtRange.PriceRange.End,
			invalidJoinedAtRange.QuantityRange.Start, invalidJoinedAtRange.QuantityRange.End, invalidJoinedAtRange.JoinedAtRange.Start.Format(time.RFC3339),
			invalidJoinedAtRange.JoinedAtRange.End.Format(time.RFC3339), invalidJoinedAtRange.DeliveredAtRange.Start.Format(time.RFC3339), invalidJoinedAtRange.DeliveredAtRange.End.Format(time.RFC3339))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid joined at range")
		assert.Empty(t, search)
	})

	t.Run("InvalidDeliveredAtRange", func(t *testing.T) {
		search, err := New(invalidDeliveredAtRange.ClientID, invalidDeliveredAtRange.Port, invalidDeliveredAtRange.Vault, invalidDeliveredAtRange.GuideNumber,
			invalidDeliveredAtRange.Type, invalidDeliveredAtRange.VehiclePlate, invalidDeliveredAtRange.PriceRange.Start, invalidDeliveredAtRange.PriceRange.End,
			invalidDeliveredAtRange.QuantityRange.Start, invalidDeliveredAtRange.QuantityRange.End, invalidDeliveredAtRange.JoinedAtRange.Start.Format(time.RFC3339),
			invalidDeliveredAtRange.JoinedAtRange.End.Format(time.RFC3339), invalidDeliveredAtRange.DeliveredAtRange.Start.Format(time.RFC3339), invalidDeliveredAtRange.DeliveredAtRange.End.Format(time.RFC3339))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid delivered at range")
		assert.Empty(t, search)
	})
}

func parseTimeIgnoringError(s string) time.Time {
	t, _ := time.Parse(time.RFC3339, s)
	return t
}
