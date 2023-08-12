package search

import (
	"fmt"
	"time"
)

// validatePriceRange checks if the start price is less than or equal to the end price.
// If not, it returns an error indicating an invalid price range.
func validatePriceRange(startPrice, endPrice float64) (err error) {
	if startPrice > endPrice {
		err = fmt.Errorf("invalid price range: start price must be less than end price")
	}
	return
}

// validateQuantityRange checks if the start quantity is less than or equal to the end quantity.
// If not, it returns an error indicating an invalid quantity range.
func validateQuantityRange(startQuantity, endQuantity int) (err error) {
	if startQuantity > endQuantity {
		err = fmt.Errorf("invalid quantity range: start quantity must be less than end quantity")
	}
	return
}

// validateDeliveredAtRange checks if the end delivered datetime is after the start delivered datetime.
// If not, it returns an error indicating an invalid delivered at range.
func validateDeliveredAtRange(startDeliveredAt, endDeliveredAt time.Time) (err error) {
	if endDeliveredAt.Before(startDeliveredAt) {
		err = fmt.Errorf("invalid delivered at range: start delivered datetime must be earlier than end delivered datetime")
	}
	return
}

// validateJoinedAtRange checks if the end joined datetime is after the start joined datetime.
// If not, it returns an error indicating an invalid joined at range.
func validateJoinedAtRange(startJoinedAt, endJoinedAt time.Time) (err error) {
	if endJoinedAt.Before(startJoinedAt) {
		err = fmt.Errorf("invalid joined at range: start joined datetime must be earlier than end joined datetime")
	}
	return
}
