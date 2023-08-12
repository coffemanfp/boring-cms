package database

import "github.com/coffemanfp/test/client"

// Constant CLIENT_REPOSITORY is used to uniquely identify the client repository.
const CLIENT_REPOSITORY RepositoryID = "CLIENT_REPOSITORYY"

// ClientRepository defines the methods for working with client data in the database.
type ClientRepository interface {
	// Get retrieves a list of clients based on the given page number.
	Get(page int) (clients []*client.Client, err error)

	// GetOne retrieves a specific client based on the provided ID.
	GetOne(id int) (client client.Client, err error)
}
