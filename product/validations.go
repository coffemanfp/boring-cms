package product

import (
	"fmt"
	"regexp"
)

func validateVehiclePlate(vp string) (err error) {
	r := regexp.MustCompile(`^[A-Za-z]{3}-[0-9]{3}$`)
	if !r.MatchString(vp) {
		err = fmt.Errorf("invalid vehicle plate: invalid vehicle plate format of %s", vp)
	}
	return
}

func validateGuideNumber(vp string) (err error) {
	r := regexp.MustCompile(`^[A-Za-z0-9]{10}$`)
	if !r.MatchString(vp) {
		err = fmt.Errorf("invalid guide number: invalid guide number format of %s", vp)
	}
	return
}

func validateCreator(createdby int) (err error) {
	if createdby <= 0 {
		err = fmt.Errorf("invalid creator id or not provided: %d", createdby)
	}
	return
}
