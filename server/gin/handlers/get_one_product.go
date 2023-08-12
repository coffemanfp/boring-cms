package handlers

import (
	"net/http"

	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/product"
	"github.com/gin-gonic/gin"
)

// GetProduct is a struct representing the action of getting a product.
type GetProduct struct{}

// Do is a method of the GetProduct struct that executes the action of getting a product.
func (gp GetProduct) Do(c *gin.Context) {
	// Retrieve the product repository.
	repo, ok := getProductRepository(c)
	if !ok {
		return
	}

	// Read the product ID from the request.
	id, ok := gp.readProductID(c)
	if !ok {
		return
	}

	// Retrieve the product from the database using the getProductFromDB method.
	p, ok := gp.getProductFromDB(c, id, repo)
	if !ok {
		return
	}

	// Generate a discount for the product using the generateDiscount method.
	p = gp.generateDiscount(p)

	// Return the product as JSON response.
	c.JSON(http.StatusOK, p)
}

// readProductID is a method of the GetProduct struct that reads and returns the product ID from the URL parameter.
func (gp GetProduct) readProductID(c *gin.Context) (id int, ok bool) {
	return readIntFromURL(c, "id", false)
}

// getProductFromDB is a method of the GetProduct struct that retrieves a product from the database.
func (gp GetProduct) getProductFromDB(c *gin.Context, id int, repo database.ProductRepository) (p product.Product, ok bool) {
	// Retrieve the product from the database using the product repository.
	p, err := repo.GetOne(id, c.GetInt("id"))
	if err != nil {
		// Handle any error and abort the request.
		handleError(c, err)
		return
	}
	ok = true
	return
}

// generateDiscount is a method of the GetProduct struct that calculates and adds a discount to the product.
func (gp GetProduct) generateDiscount(pr product.Product) (p product.Product) {
	// Clone the product to avoid modifying the original object.
	p = pr
	var vault, port int
	if p.Vault != nil {
		vault = *p.Vault
	}
	if p.Port != nil {
		port = *p.Port
	}
	// Create a discount generator based on the product's attributes.
	discountGenerator := product.NewDiscountGenerator(vault, port, *p.Quantity, *p.ShippingPrice)
	// Generate the discount using the discount generator.
	p.Discount = discountGenerator.Generate()
	return
}
