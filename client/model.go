package client

import (
	"time"

	"github.com/coffemanfp/docucentertest/auth"
	"github.com/coffemanfp/docucentertest/utils"
)

// Client represents a client entity with authentication details.
type Client struct {
	auth.Auth           // Embedding the Auth struct from the auth package.
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Surname   string    `json:"surname,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// New creates a new client using the provided clientR and performs necessary initialization.
func New(clientR Client) (client Client, err error) {
	// Clean the username by removing spaces and converting special characters.
	client.Auth.Username = utils.RemoveSpaceAndConvertSpecialChars(clientR.Auth.Username)

	// Validate the cleaned username.
	err = ValidateUsername(client.Auth.Username)
	if err != nil {
		// Reset the client to an empty state in case of an error.
		return Client{}, err
	}

	// Copy the clientR data to the new client.
	client = clientR

	// Hash the password using the HashPassword function from the auth package.
	client.Auth.Password, err = auth.HashPassword(clientR.Auth.Password)
	if err != nil {
		// Reset the client to an empty state in case of an error.
		return Client{}, err
	}

	// Clean the name by removing spaces and converting special characters.
	client.Name = utils.RemoveSpaceAndConvertSpecialChars(clientR.Name)

	// Clean the surname by removing spaces and converting special characters.
	client.Surname = utils.RemoveSpaceAndConvertSpecialChars(clientR.Surname)

	// Set the CreatedAt field to the current time.
	client.CreatedAt = time.Now()

	return
}
