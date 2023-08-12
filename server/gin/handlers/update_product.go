package handlers

import (
	"net/http"

	"github.com/coffemanfp/docucentertest/database"
	"github.com/coffemanfp/docucentertest/product"
	"github.com/coffemanfp/docucentertest/server/errors"
	"github.com/gin-gonic/gin"
)

// UpdateProduct is a struct that represents the logic for updating product information.
type UpdateProduct struct{}

// Do updates a product based on the provided data in the request.
func (up UpdateProduct) Do(c *gin.Context) {
	// Read the updated product data from the request
	t, ok := up.readProduct(c)
	if !ok {
		return
	}

	// Read the product ID from the URL parameter
	id, ok := up.readProductID(c)
	if !ok {
		return
	}

	// Update the product data based on the provided information
	t, ok = up.updateProduct(c, id, t)
	if !ok {
		return
	}

	// Retrieve the product repository
	repo, ok := getProductRepository(c)
	if !ok {
		return
	}

	// Update the product data in the database
	ok = up.updateProductInDB(c, repo, t)
	if !ok {
		return
	}

	// Respond with a success status
	c.Status(http.StatusOK)
}

func (up UpdateProduct) readProduct(c *gin.Context) (p product.Product, ok bool) {
	// Read the updated product data from the request
	ok = readRequestData(c, &p)
	return
}

func (up UpdateProduct) updateProduct(c *gin.Context, id int, pr product.Product) (p product.Product, ok bool) {
	// Update the product information using the provided data
	p, err := product.Update(pr)
	if err != nil {
		// Handle the error and return a bad request response
		err = errors.NewHTTPError(http.StatusBadRequest, err.Error())
		handleError(c, err)
		return
	}
	// Assign the ID of the updated product
	p.ID = id
	// Assign the client ID from the context
	p.ClientID = c.GetInt("id")
	ok = true
	return
}

func (up UpdateProduct) readProductID(c *gin.Context) (id int, ok bool) {
	// Read the product ID from the URL parameter
	return readIntFromURL(c, "id", false)
}

func (up UpdateProduct) updateProductInDB(c *gin.Context, repo database.ProductRepository, p product.Product) (ok bool) {
	// Update the product in the database using the provided data
	err := repo.Update(p)
	if err != nil {
		// Handle the error and return a response
		handleError(c, err)
		return
	}
	ok = true
	return
}
