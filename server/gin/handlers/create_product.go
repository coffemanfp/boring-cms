package handlers

import (
	"net/http"

	"github.com/coffemanfp/docucentertest/database"
	"github.com/coffemanfp/docucentertest/product"
	"github.com/coffemanfp/docucentertest/server/errors"
	"github.com/gin-gonic/gin"
)

// CreateProduct is a struct that represents a create product operation.
type CreateProduct struct{}

// Do is a method of the CreateProduct struct that handles the creation of a new product.
// It reads the product data from the request, creates the product, saves it in the database,
// and sends the created product back as a JSON response.
func (ct CreateProduct) Do(c *gin.Context) {
	// Read the product data from the request.
	p, ok := ct.readProduct(c)
	if !ok {
		return
	}

	// Create the product and handle any errors.
	p, ok = ct.createProduct(c, p)
	if !ok {
		return
	}

	// Get the product repository.
	repo, ok := getProductRepository(c)
	if !ok {
		return
	}

	// Save the product in the database and handle any errors.
	id, ok := ct.saveProductInDB(c, repo, p)
	if !ok {
		return
	}

	// Set the generated ID in the product.
	p.ID = id

	// Send the created product as a JSON response with a 201 Created status.
	c.JSON(http.StatusCreated, p)
}

// readProduct is a method of the CreateProduct struct that reads the product data from the request.
// It uses the readRequestData function to bind the JSON request data to a product.Product instance.
// If successful, it returns the product instance and a boolean indicating success.
func (ct CreateProduct) readProduct(c *gin.Context) (p product.Product, ok bool) {
	// Use the readRequestData function to bind JSON request data to a product.Product instance.
	ok = readRequestData(c, &p)
	return
}

// createProduct is a method of the CreateProduct struct that creates a new product based on the provided data.
// It sets the client ID from the context if not provided in the request data and validates the product data.
// If successful, it returns the created product instance and a boolean indicating success.
func (ct CreateProduct) createProduct(c *gin.Context, pr product.Product) (p product.Product, ok bool) {
	// If the client ID is not provided in the request data, use the client ID from the context.
	if pr.ClientID == 0 {
		pr.ClientID = c.GetInt("id")
	}

	// Create a new product instance based on the provided data and validate it.
	p, err := product.New(pr)
	if err != nil {
		err = errors.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		handleError(c, err)
		return
	}
	ok = true
	return
}

// saveProductInDB is a method of the CreateProduct struct that saves the created product in the database.
// It uses the provided ProductRepository to call the Create method and saves the product.
// If successful, it returns the generated ID and a boolean indicating success.
func (ct CreateProduct) saveProductInDB(c *gin.Context, repo database.ProductRepository, p product.Product) (id int, ok bool) {
	// Use the ProductRepository to create and save the product in the database.
	id, err := repo.Create(p)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
