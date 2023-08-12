package handlers

import (
	"net/http"

	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/product"
	"github.com/gin-gonic/gin"
)

type GetSomeProducts struct{}

// Do is a method of the GetSomeProducts struct that retrieves a list of products from the database,
// applies discounts to them, and sends the response in JSON format.
// It takes a gin.Context as a parameter.
func (gsp GetSomeProducts) Do(c *gin.Context) {
	// Read the page number from the request URL.
	page, ok := readPagination(c)
	if !ok {
		return
	}

	// Get the product repository.
	repo, ok := getProductRepository(c)
	if !ok {
		return
	}

	// Retrieve the list of products from the database using the specified page.
	ps, ok := gsp.getFromDB(c, repo, page)
	if !ok {
		return
	}

	// Apply discounts to the products.
	ps = gsp.generateDiscount(ps)

	// Send the list of products with applied discounts in JSON format as the response.
	c.JSON(http.StatusOK, ps)
}

// getFromDB is a method of the GetSomeProducts struct that retrieves a list of products from the database.
// It takes a gin.Context, a product repository, and a page number as parameters.
// It returns a list of products and a boolean indicating whether the operation was successful.
func (gsp GetSomeProducts) getFromDB(c *gin.Context, repo database.ProductRepository, page int) (ps []*product.Product, ok bool) {
	// Retrieve the list of products from the repository using the specified page and client ID.
	ps, err := repo.Get(page, c.GetInt("id"))
	if err != nil {
		// If there's an error, handle it and set ok to false.
		handleError(c, err)
		return
	}
	// If successful, set ok to true.
	ok = true
	return
}

// generateDiscount is a method of the GetSomeProducts struct that applies discounts to a list of products.
// It takes a list of products as a parameter and returns the same list with applied discounts.
func (gsp GetSomeProducts) generateDiscount(psR []*product.Product) (ps []*product.Product) {
	// Loop through the list of products and calculate discounts for each product.
	ps = make([]*product.Product, len(psR))
	for i, p := range psR {
		var vault, port int
		if p.Vault != nil {
			vault = *p.Vault
		}
		if p.Port != nil {
			port = *p.Port
		}
		// Create a discount generator based on the product's attributes.
		discountGenerator := product.NewDiscountGenerator(vault, port, *p.Quantity, *p.ShippingPrice)
		// Apply the discount generator to calculate the discount for the product.
		p.Discount = discountGenerator.Generate()
		// Assign the modified product to the new list of products.
		ps[i] = p
	}
	return
}
