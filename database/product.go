package database

import (
	"github.com/coffemanfp/test/product"
)

const PRODUCT_REPOSITORY RepositoryID = "PRODUCT_REPOSITORY"

type ProductRepository interface {
	Get(page int) (products []*product.Product, err error)
	GetOne(id int) (product product.Product, err error)
	Create(product product.Product) (id int, err error)
}
