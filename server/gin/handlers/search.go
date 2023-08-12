package handlers

import (
	"net/http"

	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/product"
	"github.com/coffemanfp/test/search"
	"github.com/gin-gonic/gin"
)

// Search represents a search handler for products.
type Search struct{}

// Do performs the product search based on the provided search parameters.
func (s Search) Do(c *gin.Context) {
	// Read the search parameters from the request
	srch, ok := s.readSearch(c)
	if !ok {
		return
	}

	// Get the product repository
	repo, ok := getProductRepository(c)
	if !ok {
		return
	}

	// Perform the product search in the database
	ps, ok := s.searchOnDB(c, repo, srch)
	if !ok {
		return
	}

	// Apply discount calculation to the search results
	ps = s.generateDiscount(ps)

	// Respond with the search results
	c.JSON(http.StatusOK, ps)
}
func (s Search) readSearch(c *gin.Context) (srch search.Search, ok bool) {
	// Read search parameters from query string
	guideNumber := c.Query("guideNumber")
	vehiclePlate := c.Query("vehiclePlate")
	productType := c.Query("type")
	startJoinedAt := c.Query("startJoinedAt")
	endJoinedAt := c.Query("endJoinedAt")
	startDeliveredAt := c.Query("startDeliveredAt")
	endDeliveredAt := c.Query("endDeliveredAt")
	clientID := c.GetInt("id")

	// Read port and vault parameters from URL
	port, ok := readIntFromURL(c, "port", true)
	if !ok {
		return
	}
	vault, ok := readIntFromURL(c, "vault", true)
	if !ok {
		return
	}

	// Read price range parameters from URL
	startPrice, ok := readFloatFromURL(c, "startPrice", true)
	if !ok {
		return
	}
	endPrice, ok := readFloatFromURL(c, "endPrice", true)
	if !ok {
		return
	}

	// Read quantity range parameters from URL
	startQuantity, ok := readIntFromURL(c, "startQuantity", true)
	if !ok {
		return
	}
	endQuantity, ok := readIntFromURL(c, "endQuantity", true)
	if !ok {
		return
	}

	// Create a new Search object based on the collected parameters
	srch, err := search.New(clientID, port, vault, guideNumber, productType, vehiclePlate, startPrice, endPrice, startQuantity,
		endQuantity, startJoinedAt, endJoinedAt, startDeliveredAt, endDeliveredAt)
	if err != nil {
		// Handle errors by aborting the request and sending an error response
		handleError(c, err)
		return
	}

	// Return the constructed search object and the status of the operation
	ok = true
	return
}

func (s Search) searchOnDB(c *gin.Context, repo database.ProductRepository, srch search.Search) (ps []*product.Product, ok bool) {
	// Search for products in the database based on the given search criteria
	ps, err := repo.Search(srch)
	if err != nil {
		// Handle errors by aborting the request and sending an error response
		handleError(c, err)
		return
	}
	ok = true
	return
}

func (s Search) generateDiscount(psR []*product.Product) (ps []*product.Product) {
	// Create a discount generator for each product and update the discount values
	ps = make([]*product.Product, len(psR))
	for i, p := range psR {
		var vault, port int
		if p.Vault != nil {
			vault = *p.Vault
		}
		if p.Port != nil {
			port = *p.Port
		}
		// Create a new discount generator and calculate the discount
		discountGenerator := product.NewDiscountGenerator(vault, port, *p.Quantity, *p.ShippingPrice)
		p.Discount = discountGenerator.Generate()
		ps[i] = p
	}
	return
}
