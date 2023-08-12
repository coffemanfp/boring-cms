package product

import (
	"fmt"
	"time"
)

// Product represents a product with various attributes.
type Product struct {
	ID            int        `json:"id,omitempty"`             // Unique identifier for the product.
	ClientID      int        `json:"client_id,omitempty"`      // Identifier of the associated client.
	GuideNumber   *string    `json:"guide_number,omitempty"`   // Guide number for the product, can be nil.
	Type          *string    `json:"type,omitempty"`           // Type of the product, can be nil.
	Quantity      *int       `json:"quantity,omitempty"`       // Quantity of the product, can be nil.
	JoinedAt      *time.Time `json:"joined_at,omitempty"`      // Timestamp when the product was joined, can be nil.
	DeliveredAt   *time.Time `json:"delivered_at,omitempty"`   // Timestamp when the product was delivered, can be nil.
	ShippingPrice *float64   `json:"shipping_price,omitempty"` // Shipping price of the product, can be nil.
	VehiclePlate  *string    `json:"vehicle_plate,omitempty"`  // Vehicle plate associated with the product, can be nil.
	Port          *int       `json:"port,omitempty"`           // Port associated with the product, can be nil.
	Vault         *int       `json:"vault,omitempty"`          // Vault associated with the product, can be nil.
	Discount      float64    `json:"discount,omitempty"`       // Discount applied to the product.
}

// New creates a new Product instance while validating certain fields.
func New(productR Product) (product Product, err error) {
	err = validateCreator(productR.ClientID) // Validate the associated client ID.
	if err != nil {
		return
	}

	if productR.Port != nil {
		err = ValidatePort(*productR.Port)
		if err != nil {
			return
		}
	}
	if productR.Vault != nil {
		err = ValidateVault(*productR.Vault)
		if err != nil {
			return
		}
	}

	if productR.GuideNumber == nil || *productR.GuideNumber == "" {
		err = fmt.Errorf("invalid guide number: guide number cannot be empty")
		return
	} else {
		err = ValidateGuideNumber(productR.GuideNumber)
		if err != nil {
			return
		}
	}

	err = ValidateVehiclePlate(productR.VehiclePlate)
	if err != nil {
		return
	}

	product = productR // Assign the validated product to the result.
	return
}

// Update updates a product while validating the vehicle plate and guide number.
func Update(productR Product) (product Product, err error) {
	// Check if the vehicle plate is provided and validate it.
	if productR.VehiclePlate != nil {
		err = ValidateVehiclePlate(productR.VehiclePlate)
		if err != nil {
			return // Return if there's an error in validating the vehicle plate.
		}
	}

	// Check if the guide number is provided and validate it.
	if productR.GuideNumber != nil {
		err = ValidateGuideNumber(productR.GuideNumber)
		if err != nil {
			return // Return if there's an error in validating the guide number.
		}
	}

	product = productR // If validations are successful, assign the updated product.
	return
}

// DiscountGenerator is an interface for generating discounts.
type DiscountGenerator interface {
	Generate() (discount float64) // Generate calculates and returns the discount.
}

// DiscountGeneratorImpl is an implementation of the DiscountGenerator interface.
type DiscountGeneratorImpl struct {
	quantity      int     // The quantity of products.
	vault         int     // The number of vaults.
	port          int     // The number of ports.
	shippingPrice float64 // The shipping price.
}

// Generate calculates the discount based on certain conditions.
func (dgi DiscountGeneratorImpl) Generate() (discount float64) {
	// Check if the quantity is less than 10.
	if dgi.quantity < 10 {
		return // Return without applying any discount.
	}

	// Check if there are vaults or ports for potential discounts.
	if dgi.vault > 0 {
		discount = dgi.shippingPrice * 0.05 // Apply a discount of 5%.
	} else if dgi.port > 0 {
		discount = dgi.shippingPrice * 0.03 // Apply a discount of 3%.
	}

	return // Return the calculated discount.
}

// NewDiscountGenerator creates a new DiscountGenerator instance.
func NewDiscountGenerator(vault, port, quantity int, shippingPrice float64) DiscountGenerator {
	return DiscountGeneratorImpl{
		vault:         vault,
		quantity:      quantity,
		port:          port,
		shippingPrice: shippingPrice,
	}
}
