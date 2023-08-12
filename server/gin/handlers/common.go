package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/server/errors"
	"github.com/gin-gonic/gin"
)

// readRequestData takes a Gin context (c) and a struct (v), and tries to read and bind JSON data from the request into the provided struct.
// If successful, it returns ok as true, otherwise, it handles the error and returns ok as false.
func readRequestData(c *gin.Context, v interface{}) (ok bool) {
	// Try to bind JSON data from the request to the provided struct.
	err := c.ShouldBindJSON(v)
	if err != nil {
		// If there's an error during binding, create an HTTP error and handle it using the handleError function.
		err = errors.NewHTTPError(http.StatusBadRequest, err.Error())
		handleError(c, err)
		return
	}
	// Indicate that reading and binding were successful.
	ok = true
	return
}

// getAuthRepository tries to retrieve an instance of the AuthRepository from the repository map.
// If successful, it returns the retrieved repository and ok as true. If there's an error, it handles the error and returns ok as false.
func getAuthRepository(c *gin.Context) (repo database.AuthRepository, ok bool) {
	// Attempt to get the AuthRepository from the repository map using the appropriate key.
	repo, err := database.GetRepository[database.AuthRepository](db, database.AUTH_REPOSITORY)
	if err != nil {
		// If there's an error while retrieving the repository, handle the error using the handleError function.
		handleError(c, err)
		return
	}
	// Indicate that the repository retrieval was successful.
	ok = true
	return
}

// getClientRepository tries to retrieve an instance of the ClientRepository from the repository map.
// If successful, it returns the retrieved repository and ok as true. If there's an error, it handles the error and returns ok as false.
func getClientRepository(c *gin.Context) (repo database.ClientRepository, ok bool) {
	// Attempt to get the ClientRepository from the repository map using the appropriate key.
	repo, err := database.GetRepository[database.ClientRepository](db, database.CLIENT_REPOSITORY)
	if err != nil {
		// If there's an error while retrieving the repository, handle the error using the handleError function.
		handleError(c, err)
		return
	}
	// Indicate that the repository retrieval was successful.
	ok = true
	return
}

// getProductRepository tries to retrieve an instance of the ProductRepository from the repository map.
// If successful, it returns the retrieved repository and ok as true. If there's an error, it handles the error and returns ok as false.
func getProductRepository(c *gin.Context) (repo database.ProductRepository, ok bool) {
	repo, err := database.GetRepository[database.ProductRepository](db, database.PRODUCT_REPOSITORY)
	if err != nil {
		// If there's an error while retrieving the repository, handle the error using the handleError function.
		handleError(c, err)
		return
	}
	// Indicate that the repository retrieval was successful.
	ok = true
	return
}

// readIntFromURL reads an integer value from the URL parameter or query parameter based on isQueryParam.
// It returns the parsed integer value and ok as true if successful. If the parameter is empty, it returns ok as true without value.
// If parsing fails or the parameter is invalid, it creates an HTTP error and handles it using the handleError function, returning ok as false.
func readIntFromURL(c *gin.Context, param string, isQueryParam bool) (v int, ok bool) {
	// Get the parameter value from the URL based on whether it's a query parameter or not.
	var p string
	if isQueryParam {
		p = c.Query(param)
	} else {
		p = c.Param(param)
	}
	// If the parameter is empty, return without an error.
	if p == "" {
		ok = true
		return
	}
	// Parse the parameter value as an integer.
	v, err := strconv.Atoi(p)
	if err != nil {
		// If parsing fails, create an HTTP error and handle it using the handleError function.
		err = fmt.Errorf("invalid %s param: %s", param, p)
		err = errors.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		handleError(c, err)
		return
	}

	// Indicate that the parameter parsing was successful.
	ok = true
	return
}

// readFloatFromURL reads a floating-point value from the URL parameter or query parameter based on isQueryParam.
// It returns the parsed floating-point value and ok as true if successful. If the parameter is empty, it returns ok as true without value.
// If parsing fails or the parameter is invalid, it creates an HTTP error and handles it using the handleError function, returning ok as false.
func readFloatFromURL(c *gin.Context, param string, isQueryParam bool) (v float64, ok bool) {
	// Get the parameter value from the URL based on whether it's a query parameter or not.
	var p string
	if isQueryParam {
		p = c.Query(param)
	} else {
		p = c.Param(param)
	}
	// If the parameter is empty, return without an error.
	if p == "" {
		ok = true
		return
	}
	// Parse the parameter value as a floating-point number.
	v, err := strconv.ParseFloat(p, 64)
	if err != nil {
		// If parsing fails, create an HTTP error and handle it using the handleError function.
		err = fmt.Errorf("invalid %s param: %s", param, p)
		err = errors.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		handleError(c, err)
		return
	}

	// Indicate that the parameter parsing was successful.
	ok = true
	return
}

// readPagination reads the "page" parameter from the URL using readIntFromURL and returns it.
// It delegates to readIntFromURL to handle parsing and potential errors.
func readPagination(c *gin.Context) (page int, ok bool) {
	page, ok = readIntFromURL(c, "page", false)
	return
}

// handleError handles an error by logging it to the context and aborting the request.
func handleError(c *gin.Context, err error) {
	c.Error(err)
	c.Abort()
}
