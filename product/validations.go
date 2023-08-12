package product

import (
	"fmt"
	"regexp"
)

// ValidateVehiclePlate validates the format of a vehicle plate.
func ValidateVehiclePlate(vp *string) (err error) {
	r := regexp.MustCompile(`^[A-Za-z]{3}-[0-9]{3}$`) // Regular expression to match the expected format.
	if !r.MatchString(*vp) {
		err = fmt.Errorf("invalid vehicle plate: invalid vehicle plate format of %s", *vp) // If the format doesn't match, create an error.
	}
	return
}

// ValidateGuideNumber validates the format of a guide number.
func ValidateGuideNumber(gn *string) (err error) {
	r := regexp.MustCompile(`^[A-Za-z0-9]{10}$`) // Regular expression to match the expected format.
	if !r.MatchString(*gn) {
		err = fmt.Errorf("invalid guide number: invalid guide number format of %s", *gn) // If the format doesn't match, create an error.
	}
	return
}

func ValidatePort(port int) (err error) {
	if port < 0 {
		err = fmt.Errorf("invalid port: port must be a positive number %d", port)
	}
	return
}

func ValidateVault(vault int) (err error) {
	if vault < 0 {
		err = fmt.Errorf("invalid vault: vault must be a positive number %d", vault)
	}
	return
}

// validateCreator validates the ID of the creator.
func validateCreator(createdby int) (err error) {
	if createdby <= 0 {
		err = fmt.Errorf("invalid creator id or not provided: %d", createdby) // If the ID is not valid, create an error.
	}
	return
}
