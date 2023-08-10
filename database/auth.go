package database

import (
	"github.com/coffemanfp/test/auth"
	"github.com/coffemanfp/test/client"
)

// AUTH_REPOSITORY is the key to be used when creating the repositories hashmap.
const AUTH_REPOSITORY RepositoryID = "AUTH"

// AuthRepository defines the behaviors to be used by a AuthRepository implementation.
type AuthRepository interface {
	GetIdAndHashedPassword(auth auth.Auth) (id int, hash string, err error)
	Register(client client.Client) (id int, err error)
}
