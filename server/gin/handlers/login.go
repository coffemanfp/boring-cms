package handlers

import (
	"net/http"

	"github.com/coffemanfp/test/auth"
	"github.com/coffemanfp/test/client"
	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/server/errors"
	"github.com/gin-gonic/gin"
)

// Login represents the login handler responsible for authenticating users.
type Login struct{}

// Do performs the login process. It reads the client's credentials from the request,
// searches for the credentials in the authentication repository, compares the provided
// password with the hashed password from the database, generates an authentication token,
// and responds with the generated token.
func (l Login) Do(c *gin.Context) {
	// Read the client credentials from the request data
	client, ok := l.readCredentials(c)
	if !ok {
		return
	}

	// Get the authentication repository to perform database operations
	repo, ok := l.getAuthRepository(c)
	if !ok {
		return
	}

	// Search for the client's credentials in the database and retrieve the client's ID and hashed password
	id, hash, ok := l.searchCredentialsInDB(c, client, repo)
	if !ok {
		return
	}

	// Compare the provided password with the hashed password from the database
	ok = l.comparePassword(c, hash, client.Auth.Password)
	if !ok {
		return
	}

	// Generate an authentication token using the retrieved client ID
	token, ok := l.generateToken(c, id)
	if !ok {
		return
	}

	// Respond with the generated token
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// readCredentials reads and parses the client credentials from the request data.
// If successful, it returns the parsed client and true.
func (l Login) readCredentials(c *gin.Context) (client client.Client, ok bool) {
	// Read client credentials from the request data
	ok = readRequestData(c, &client)
	return
}

// getAuthRepository retrieves the authentication repository.
// If successful, it returns the authentication repository and true.
func (l Login) getAuthRepository(c *gin.Context) (repo database.AuthRepository, ok bool) {
	// Retrieve the authentication repository
	return getAuthRepository(c)
}

// searchCredentialsInDB searches for client credentials in the database
// using the provided authentication repository and client data.
// If successful, it returns the client's ID, hashed password, and true.
// If the credentials are invalid, it returns an unauthorized error and false.
func (l Login) searchCredentialsInDB(c *gin.Context, client client.Client, repo database.AuthRepository) (id int, hash string, ok bool) {
	// Search for client credentials in the database and retrieve the client's ID and hashed password
	id, hash, err := repo.GetIdAndHashedPassword(client.Auth)
	if err != nil {
		// Return an unauthorized error if the credentials are invalid
		err = errors.NewHTTPError(http.StatusUnauthorized, errors.UNAUTHORIZED_ERROR_MESSAGE)
		handleError(c, err)
		return
	}
	ok = true
	return
}

// comparePassword compares the provided password with the hashed password.
// If the comparison fails, it handles the error and returns false.
// If the comparison succeeds, it returns true.
func (l Login) comparePassword(c *gin.Context, hash, password string) (ok bool) {
	// Compare the provided password with the hashed password
	err := auth.CompareHashAndPassword(hash, password)
	if err != nil {
		// Handle error if password comparison fails
		handleError(c, err)
		return
	}
	ok = true
	return
}

// generateToken generates a JWT token using the client's ID, JWT lifespan,
// and secret key from the configuration settings.
// If token generation fails, it handles the error and returns false.
// If token generation succeeds, it returns the generated token and true.
func (l Login) generateToken(c *gin.Context, id int) (token string, ok bool) {
	// Generate a JWT token using the client's ID and configuration settings
	token, err := auth.GenerateToken(id, conf.Server.JWTLifespan, conf.Server.SecretKey)
	if err != nil {
		// Handle error if token generation fails
		handleError(c, err)
		return
	}
	ok = true
	return
}
