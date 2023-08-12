package handlers

import (
	"net/http"

	"github.com/coffemanfp/test/client"
	"github.com/coffemanfp/test/database"
	"github.com/gin-gonic/gin"
)

// GetClient represents a struct for handling the action of retrieving a client.
type GetClient struct{}

// Do is a method of the GetClient struct that executes the action of retrieving a client.
func (gc GetClient) Do(c *gin.Context) {
	// Read the client ID from the request using the readClientID method.
	id, ok := gc.readClientID(c)
	if !ok {
		return
	}

	// Get the client repository using the getClientRepository method.
	repo, ok := getClientRepository(c)
	if !ok {
		return
	}

	// Retrieve the client from the database using the getClientFromDB method.
	cl, ok := gc.getClientFromDB(c, repo, id)
	if !ok {
		return
	}

	// Respond with the retrieved client in JSON format.
	c.JSON(http.StatusOK, cl)
}

// readClientID is a method of the GetClient struct that reads the client ID from the request.
// It returns the parsed client ID and a boolean indicating if the operation was successful.
func (gc GetClient) readClientID(c *gin.Context) (id int, ok bool) {
	// Use the readIntFromURL function to read the client ID from the URL parameters.
	return readIntFromURL(c, "id", false)
}

// getClientFromDB is a method of the GetClient struct that retrieves a client from the database.
// It returns the retrieved client and a boolean indicating if the operation was successful.
func (gc GetClient) getClientFromDB(c *gin.Context, repo database.ClientRepository, id int) (cl client.Client, ok bool) {
	// Retrieve the client using the GetOne method of the client repository.
	cl, err := repo.GetOne(id)
	if err != nil {
		// If an error occurs, handle it and return false.
		handleError(c, err)
		return
	}
	// Return true to indicate a successful operation.
	ok = true
	return
}
