package database

import (
	"github.com/coffemanfp/test/product"
	"github.com/coffemanfp/test/search"
)

const PRODUCT_REPOSITORY RepositoryID = "PRODUCT_REPOSITORY"

type ProductRepository interface {
	Get(page, clientID int) (products []*product.Product, err error)
	GetOne(id, clientID int) (product product.Product, err error)
	Create(product product.Product) (id int, err error)
	Search(search search.Search) (products []*product.Product, err error)
	Update(product product.Product) (err error)
	Delete(id, clientID int) (err error)
}
