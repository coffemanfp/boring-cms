package product

import "time"

type Product struct {
	ID            int       `json:"id,omitempty"`
	ClientID      int       `json:"client_id,omitempty"`
	GuideNumber   string    `json:"guide_number,omitempty"`
	Type          string    `json:"type,omitempty"`
	JoinedAt      time.Time `json:"joined_at,omitempty"`
	DeliveredAt   time.Time `json:"delivered_at,omitempty"`
	ShippingPrice float64   `json:"shipping_price,omitempty"`
	VehiclePlate  string    `json:"vehicle_plate,omitempty"`
	Port          int       `json:"port,omitempty"`
	Vault         int       `json:"vault,omitempty"`
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

	product = productR
	return
}
