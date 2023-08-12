package product

import "time"

type Product struct {
	ID            int        `json:"id,omitempty"`
	ClientID      int        `json:"client_id,omitempty"`
	GuideNumber   *string    `json:"guide_number,omitempty"`
	Type          *string    `json:"type,omitempty"`
	Quantity      *int       `json:"quantity,omitempty"`
	JoinedAt      *time.Time `json:"joined_at,omitempty"`
	DeliveredAt   *time.Time `json:"delivered_at,omitempty"`
	ShippingPrice *float64   `json:"shipping_price,omitempty"`
	VehiclePlate  *string    `json:"vehicle_plate,omitempty"`
	Port          *int       `json:"port,omitempty"`
	Vault         *int       `json:"vault,omitempty"`
	Discount      float64    `json:"discount,omitempty"`
}

func New(productR Product) (product Product, err error) {
	err = validateCreator(productR.ClientID)
	if err != nil {
		return
	}

	err = validateVehiclePlate(productR.VehiclePlate)
	if err != nil {
		return
	}

	err = validateGuideNumber(productR.GuideNumber)
	if err != nil {
		return
	}

	product = productR
	return
}

func Update(productR Product) (product Product, err error) {
	if productR.VehiclePlate != nil {
		err = validateVehiclePlate(productR.VehiclePlate)
		if err != nil {
			return
		}
	}
	if productR.GuideNumber != nil {
		err = validateGuideNumber(productR.GuideNumber)
		if err != nil {
			return
		}
	}
	product = productR
	return
}

type DiscountGenerator interface {
	Generate() (discount float64)
}

type DiscountGeneratorImpl struct {
	quantity      int
	vault         int
	port          int
	shippingPrice float64
}

func (dgi DiscountGeneratorImpl) Generate() (discount float64) {
	if dgi.quantity < 10 {
		return
	}

	if dgi.vault > 0 {
		discount = dgi.shippingPrice * 0.05
	} else if dgi.port > 0 {
		discount = dgi.shippingPrice * 0.03
	}
	return
}

func NewDiscountGenerator(vault, port, quantity int, shippingPrice float64) DiscountGenerator {
	return DiscountGeneratorImpl{
		vault:         vault,
		quantity:      quantity,
		port:          port,
		shippingPrice: shippingPrice,
	}
}
