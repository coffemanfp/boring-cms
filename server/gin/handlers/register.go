package handlers

import (
	"net/http"

	"github.com/coffemanfp/test/auth"
	"github.com/coffemanfp/test/client"
	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/server/errors"
	"github.com/gin-gonic/gin"
)

// Register represents the registration handler for new client accounts.
type Register struct{}

// Do performs the registration process. It reads the new client's data from the request,
// creates a new client instance, registers the client in the database, generates an authentication token,
// and responds with the generated token.
func (r Register) Do(c *gin.Context) {
	// Read the new client's data from the request
	client, ok := r.readClient(c)
	if !ok {
		return
	}

	// Create a new client instance and validate its data
	client, ok = r.createNewClient(c, client)
	if !ok {
		return
	}

	// Get the authentication repository to perform database operations
	repo, ok := r.getAuthRepository(c)
	if !ok {
		return
	}

	// Register the new client in the database and retrieve the generated client ID
	id, ok := r.registerClientInDB(c, client, repo)
	if !ok {
		return
	}

	// Generate an authentication token using the generated client ID
	token, ok := r.generateToken(c, id)
	if !ok {
		return
	}

	// Respond with the generated token and a status indicating successful account creation
	c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}

// readClient reads and validates the data of the new client from the request body.
func (r Register) readClient(c *gin.Context) (client client.Client, ok bool) {
	ok = readRequestData(c, &client)
	return
}

// createNewClient creates a new client instance and validates its data.
func (r Register) createNewClient(c *gin.Context, clientR client.Client) (cl client.Client, ok bool) {
	cl, err := client.New(clientR)
	if err != nil {
		err = errors.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		handleError(c, err)
		return
	}
	ok = true
	return
}

// getAuthRepository retrieves the authentication repository for database operations.
func (r Register) getAuthRepository(c *gin.Context) (repo database.AuthRepository, ok bool) {
	return getAuthRepository(c)
}

// registerClientInDB registers the new client in the database and retrieves the assigned ID.
func (r Register) registerClientInDB(c *gin.Context, client client.Client, repo database.AuthRepository) (id int, ok bool) {
	id, err := repo.Register(client)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}

// generateToken generates an authentication token for the registered client's ID.
func (r Register) generateToken(c *gin.Context, id int) (token string, ok bool) {
	token, err := auth.GenerateToken(id, conf.Server.JWTLifespan, conf.Server.SecretKey)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
