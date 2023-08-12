package handlers

import (
	"net/http"

	"github.com/coffemanfp/test/database"
	"github.com/gin-gonic/gin"
)

// DeleteProduct represents the action of deleting a product.
type DeleteProduct struct{}

// Do is a method of the DeleteProduct struct that performs the deletion of a product.
// It reads the product ID from the request, retrieves the product repository,
// and then deletes the product from the database using the deleteProductInDB function.
// If successful, it responds with a 200 OK status.
func (dp DeleteProduct) Do(c *gin.Context) {
	// Read the product ID from the request.
	id, ok := dp.readProductID(c)
	if !ok {
		return
	}

	// Get the product repository using the getProductRepository function.
	repo, ok := getProductRepository(c)
	if !ok {
		return
	}

	// Delete the product in the database using the deleteProductInDB function.
	ok = dp.deleteProductInDB(c, repo, id)
	if !ok {
		return
	}

	// Respond with a 200 OK status.
	c.Status(http.StatusOK)
}

// readProductID is a method of the DeleteProduct struct that reads the product ID from the URL parameter.
func (dp DeleteProduct) readProductID(c *gin.Context) (id int, ok bool) {
	// Read the product ID from the URL parameter using readIntFromURL function.
	return readIntFromURL(c, "id", false)
}

// deleteProductInDB is a method of the DeleteProduct struct that deletes a product from the database.
// It takes the product ID and product repository as parameters and uses the Delete method of the repository.
// If successful, it returns true, otherwise, it handles the error and returns false.
func (dp DeleteProduct) deleteProductInDB(c *gin.Context, repo database.ProductRepository, id int) (ok bool) {
	// Delete the product in the database using the Delete method of the repository.
	err := repo.Delete(id, c.GetInt("id"))
	if err != nil {
		// Handle the error using the handleError function.
		handleError(c, err)
		return
	}
	ok = true
	return
}
