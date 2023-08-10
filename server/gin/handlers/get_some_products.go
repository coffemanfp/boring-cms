package handlers

import (
	"net/http"

	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/product"
	"github.com/gin-gonic/gin"
)

type GetSomeProducts struct{}

func (gsp GetSomeProducts) Do(c *gin.Context) {
	page, ok := readPagination(c)
	if !ok {
		return
	}

	repo, ok := getProductRepository(c)
	if !ok {
		return
	}

	ls, ok := gsp.getFromDB(c, repo, page)
	if !ok {
		return
	}

	c.JSON(http.StatusOK, ls)
}

func (gsp GetSomeProducts) getFromDB(c *gin.Context, repo database.ProductRepository, page int) (ps []*product.Product, ok bool) {
	ps, err := repo.Get(page)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
