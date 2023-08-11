package search

import "time"

type Search struct {
	GuideNumber      string
	Type             string
	Port             int
	Vault            int
	VehiclePlate     string
	PriceRange       RangeFloat64
	QuantityRange    RangeInt
	JoinedAtRange    RangeTime
	DeliveredAtRange RangeTime
}

type RangeFloat64 struct {
	Start float64
	End   float64
}

type RangeInt struct {
	Start int
	End   int
}

type RangeTime struct {
	Start time.Time
	End   time.Time
}

func New(sr Search) (s Search) {
	return
}
