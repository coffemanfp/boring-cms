package handlers

import (
	"net/http"

	"github.com/coffemanfp/test/database"
	"github.com/gin-gonic/gin"
)

type DeleteProduct struct{}

func (dp DeleteProduct) Do(c *gin.Context) {
	id, ok := dp.readProductID(c)
	if !ok {
		return
	}

	repo, ok := getProductRepository(c)
	if !ok {
		return
	}

	ok = dp.deleteProductInDB(c, repo, id)
	if !ok {
		return
	}

	c.Status(http.StatusOK)
}

func (dp DeleteProduct) readProductID(c *gin.Context) (id int, ok bool) {
	return readIntFromURL(c, "id", false)
}

func (dp DeleteProduct) deleteProductInDB(c *gin.Context, repo database.ProductRepository, id int) (ok bool) {
	err := repo.Delete(id, c.GetInt("id"))
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
