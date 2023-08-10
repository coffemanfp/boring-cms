package handlers

import (
	"net/http"

	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/product"
	"github.com/gin-gonic/gin"
)

type GetProduct struct{}

func (gl GetProduct) Do(c *gin.Context) {
	repo, ok := getProductRepository(c)
	if !ok {
		return
	}

	id, ok := gl.readProductID(c)
	if !ok {
		return
	}

	p, ok := gl.getProductFromDB(c, id, repo)
	if !ok {
		return
	}

	c.JSON(http.StatusOK, p)
}

func (gp GetProduct) readProductID(c *gin.Context) (id int, ok bool) {
	return readIntParam(c, "id")
}

func (gp GetProduct) getProductFromDB(c *gin.Context, id int, repo database.ProductRepository) (p product.Product, ok bool) {
	p, err := repo.GetOne(id)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
