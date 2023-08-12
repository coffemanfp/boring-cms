package handlers

import (
	"net/http"

	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/product"
	"github.com/coffemanfp/test/server/errors"
	"github.com/gin-gonic/gin"
)

type UpdateProduct struct{}

func (up UpdateProduct) Do(c *gin.Context) {
	t, ok := up.readProduct(c)
	if !ok {
		return
	}

	id, ok := up.readProductID(c)
	if !ok {
		return
	}

	t, ok = up.updateProduct(c, id, t)
	if !ok {
		return
	}

	repo, ok := getProductRepository(c)
	if !ok {
		return
	}

	ok = up.updateProductInDB(c, repo, t)
	if !ok {
		return
	}

	c.Status(http.StatusOK)
}

func (up UpdateProduct) readProduct(c *gin.Context) (p product.Product, ok bool) {
	ok = readRequestData(c, &p)
	return
}

func (up UpdateProduct) updateProduct(c *gin.Context, id int, pr product.Product) (p product.Product, ok bool) {
	p, err := product.Update(pr)
	if err != nil {
		err = errors.NewHTTPError(http.StatusBadRequest, err.Error())
		handleError(c, err)
		return
	}
	p.ID = id
	p.ClientID = c.GetInt("id")
	ok = true
	return
}

func (up UpdateProduct) readProductID(c *gin.Context) (id int, ok bool) {
	return readIntFromURL(c, "id", false)
}

func (up UpdateProduct) updateProductInDB(c *gin.Context, repo database.ProductRepository, p product.Product) (ok bool) {
	err := repo.Update(p)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
