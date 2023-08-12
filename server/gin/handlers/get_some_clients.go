package handlers

import (
	"net/http"

	"github.com/coffemanfp/docucentertest/client"
	"github.com/coffemanfp/docucentertest/database"
	"github.com/gin-gonic/gin"
)

// GetSomeClients is a struct representing the action to retrieve a list of clients.
type GetSomeClients struct{}

// Do is the method of the GetSomeClients struct that performs the action.
func (gst GetSomeClients) Do(c *gin.Context) {
	// Read the page parameter from the URL.
	page, ok := readPagination(c)
	if !ok {
		return
	}

	// Get the client repository.
	repo, ok := getClientRepository(c)
	if !ok {
		return
	}

	// Retrieve the list of clients using the repository and the specified page.
	cs, ok := gst.get(c, repo, page)
	if !ok {
		return
	}

	// Return the list of clients as a JSON response.
	c.JSON(http.StatusOK, cs)
}

// get is a method of the GetSomeClients struct that retrieves a list of clients from the database.
// It takes a gin.Context, a client repository, and a page number as parameters.
// It returns a list of clients and a boolean indicating whether the operation was successful.
func (gst GetSomeClients) get(c *gin.Context, repo database.ClientRepository, page int) (cs []*client.Client, ok bool) {
	// Retrieve the list of clients from the repository using the specified page.
	cs, err := repo.Get(page)
	if err != nil {
		// If there's an error, handle it and set ok to false.
		handleError(c, err)
		return
	}
	// If successful, set ok to true.
	ok = true
	return
}
