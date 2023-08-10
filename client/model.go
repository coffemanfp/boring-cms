package client

import (
	"time"

	"github.com/coffemanfp/test/auth"
	"github.com/coffemanfp/test/utils"
)

type Client struct {
	auth.Auth
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Surname   string    `json:"surname,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func New(clientR Client) (client Client, err error) {
	client.Auth.Username = utils.RemoveSpaceAndConvertSpecialChars(clientR.Auth.Username)
	err = ValidateUsername(client.Auth.Username)
	if err != nil {
		return
	}

	client = clientR

	client.Auth.Password, err = auth.HashPassword(clientR.Auth.Password)
	if err != nil {
		client = Client{}
		return
	}
	client.Name = utils.RemoveSpaceAndConvertSpecialChars(clientR.Name)
	client.Surname = utils.RemoveSpaceAndConvertSpecialChars(clientR.Surname)
	client.CreatedAt = time.Now()
	return
}
