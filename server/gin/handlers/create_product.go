package handlers

import (
	"net/http"

	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/product"
	"github.com/coffemanfp/test/server/errors"
	"github.com/gin-gonic/gin"
)

type CreateProduct struct{}

func (ct CreateProduct) Do(c *gin.Context) {
	p, ok := ct.readProduct(c)
	if !ok {
		return
	}

	p, ok = ct.createProduct(c, p)
	if !ok {
		return
	}

	repo, ok := getProductRepository(c)
	if !ok {
		return
	}

	id, ok := ct.saveProductInDB(c, repo, p)
	if !ok {
		return
	}

	p.ID = id

	c.JSON(http.StatusCreated, p)
}

func (ct CreateProduct) readProduct(c *gin.Context) (p product.Product, ok bool) {
	ok = readRequestData(c, &p)
	return
}

func (ct CreateProduct) createProduct(c *gin.Context, pr product.Product) (p product.Product, ok bool) {
	if pr.ClientID == 0 {
		pr.ClientID = c.GetInt("id")
	}

	p, err := product.New(pr)
	if err != nil {
		err = errors.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		handleError(c, err)
		return
	}
	ok = true
	return
}

func (ct CreateProduct) saveProductInDB(c *gin.Context, repo database.ProductRepository, p product.Product) (id int, ok bool) {
	id, err := repo.Create(p)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
