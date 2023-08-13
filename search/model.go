package search

import (
	"fmt"
	"time"

	"github.com/coffemanfp/docucentertest/product"
)

// Search represents the search criteria for filtering products.
type Search struct {
	ClientID         int          // Client ID to filter products by.
	GuideNumber      string       // Guide number to filter products by.
	Type             string       // Product type to filter products by.
	Port             int          // Port number to filter products by.
	Vault            int          // Vault number to filter products by.
	VehiclePlate     string       // Vehicle plate to filter products by.
	PriceRange       RangeFloat64 // Price range to filter products by.
	QuantityRange    RangeInt     // Quantity range to filter products by.
	JoinedAtRange    RangeTime    // JoinedAt (timestamp) range to filter products by.
	DeliveredAtRange RangeTime    // DeliveredAt (timestamp) range to filter products by.
}

// RangeFloat64 represents a range of floating-point numbers.
type RangeFloat64 struct {
	Start float64 // Start of the range.
	End   float64 // End of the range.
}

// RangeInt represents a range of integer numbers.
type RangeInt struct {
	Start int // Start of the range.
	End   int // End of the range.
}

// RangeTime represents a range of time values.
type RangeTime struct {
	Start time.Time // Start of the range.
	End   time.Time // End of the range.
}

// New creates a new Search instance with the provided search criteria.
func New(clientID, port, vault int, guideNumber, productType, vehiclePlate string,
	startPrice, endPrice float64, startQuantity, endQuantity int, startJoinedAt, endJoinedAt,
	startDeliveredAt, endDeliveredAt string) (s Search, err error) {

	// Validate and set port.
	err = product.ValidatePort(port)
	if err != nil {
		return
	}

	// Validate and set vault.
	err = product.ValidateVault(vault)
	if err != nil {
		return
	}

	if guideNumber != "" {
		// Validate and set guide number.
		err = product.ValidateGuideNumber(&guideNumber)
		if err != nil {
			return
		}
	}

	if vehiclePlate != "" {
		// Validate and set vehicle plate.
		err = product.ValidateVehiclePlate(&vehiclePlate)
		if err != nil {
			return
		}
	}

	// Validate and set price range.
	err = validatePriceRange(startPrice, endPrice)
	if err != nil {
		return
	}

	// Validate and set quantity range.
	err = validateQuantityRange(startQuantity, endQuantity)
	if err != nil {
		return
	}

	// Parse and set the start and end joined at values.
	startJoinedAtTime, err := parseTimeValue(startJoinedAt, "start joined at")
	if err != nil {
		return
	}
	endJoinedAtTime, err := parseTimeValue(endJoinedAt, "end joined at")
	if err != nil {
		return
	}
	// Validate the joined at range.
	err = validateJoinedAtRange(startJoinedAtTime, endJoinedAtTime)
	if err != nil {
		return
	}

	// Parse and set the start and end delivered at values.
	startDeliveredAtTime, err := parseTimeValue(startDeliveredAt, "start delivered at")
	if err != nil {
		return
	}
	endDeliveredAtTime, err := parseTimeValue(endDeliveredAt, "end delivered at")
	if err != nil {
		return
	}
	// Validate the delivered at range.
	err = validateDeliveredAtRange(startDeliveredAtTime, endDeliveredAtTime)
	if err != nil {
		return
	}

	s.ClientID = clientID
	s.Port = port
	s.Type = productType
	s.Vault = vault
	s.GuideNumber = guideNumber
	s.VehiclePlate = vehiclePlate
	s.PriceRange.Start = startPrice
	s.PriceRange.End = endPrice
	s.QuantityRange.Start = startQuantity
	s.QuantityRange.End = endQuantity
	s.JoinedAtRange.Start = startJoinedAtTime
	s.JoinedAtRange.End = endJoinedAtTime
	s.DeliveredAtRange.Start = startDeliveredAtTime
	s.DeliveredAtRange.End = endDeliveredAtTime
	return
}

// parseTimeValue parses a string value to a time.Time instance.
// If the input value is empty, it returns a zero time value.
// Otherwise, it attempts to parse the input value using RFC3339 format.
// If parsing fails, it returns an error.
func parseTimeValue(v string, name string) (t time.Time, err error) {
	// If the input value is empty, return zero time value.
	if v == "" {
		return
	}
	// Attempt to parse the input value using RFC3339 format.
	t, err = time.Parse(time.RFC3339, v)
	if err != nil {
		err = fmt.Errorf("invalid %s: invalid %s time format of %s", name, name, v)
	}
	return
}
