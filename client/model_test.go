package client

import (
	"testing"

	"github.com/coffemanfp/docucentertest/auth"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestNewClient(t *testing.T) {
	clientR := Client{
		ID: 5,
		Auth: auth.Auth{
			Username: "testuser",
			Password: "testpassword",
		},
		Name:    "John",
		Surname: "Doe",
	}

	client, err := New(clientR)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.NotEmpty(t, client.ID)
	assert.NotEmpty(t, client.CreatedAt)

	// Verify the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(client.Auth.Password), []byte("testpassword"))
	assert.NoError(t, err)

	// Verify cleaned and validated username
	assert.Equal(t, "testuser", client.Auth.Username)

	// Verify cleaned name and surname
	assert.Equal(t, "John", client.Name)
	assert.Equal(t, "Doe", client.Surname)
}

func TestNewClient_InvalidUsername(t *testing.T) {
	clientR := Client{
		Auth: auth.Auth{
			Username: "inv@lidUser",
			Password: "testpassword",
		},
		Name:    "John",
		Surname: "Doe",
	}

	_, err := New(clientR)
	assert.Error(t, err)
}

func TestNewClient_InvalidNameOrSurname(t *testing.T) {
	clientR := Client{
		Auth: auth.Auth{
			Username: "inv@liduser",
			Password: "testpassword",
		},
		Name:    "Name",
		Surname: "Surname",
	}

	client, err := New(clientR)
	assert.Error(t, err)
	assert.Empty(t, client)
}
