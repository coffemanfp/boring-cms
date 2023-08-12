package database

import (
	"github.com/coffemanfp/test/product"
	"github.com/coffemanfp/test/search"
)

// Constant PRODUCT_REPOSITORY is used to uniquely identify the product repository.
const PRODUCT_REPOSITORY RepositoryID = "PRODUCT_REPOSITORY"

// ProductRepository defines the methods for working with product data in the database.
type ProductRepository interface {
	// Get retrieves a list of products based on the given page and client ID.
	Get(page, clientID int) (products []*product.Product, err error)

	// GetOne retrieves a specific product based on the provided ID and client ID.
	GetOne(id, clientID int) (product product.Product, err error)

	// Create inserts a new product into the database and returns its ID.
	Create(product product.Product) (id int, err error)

	// Search retrieves a list of products based on the provided search criteria.
	Search(search search.Search) (products []*product.Product, err error)

	// Update updates the details of a product in the database.
	Update(product product.Product) (err error)

	// Delete removes a product from the database based on the provided ID and client ID.
	Delete(id, clientID int) (err error)
}
