package handlers

import (
	"net/http"

	"github.com/coffemanfp/test/client"
	"github.com/coffemanfp/test/database"
	"github.com/gin-gonic/gin"
)

type GetClient struct{}

func (gc GetClient) Do(c *gin.Context) {
	id, ok := gc.readClientID(c)
	if !ok {
		return
	}

	repo, ok := getClientRepository(c)
	if !ok {
		return
	}

	cl, ok := gc.getClientFromDB(c, repo, id)
	if !ok {
		return
	}

	c.JSON(http.StatusOK, cl)
}

func (gc GetClient) readClientID(c *gin.Context) (id int, ok bool) {
	return readIntParam(c, "id")
}

func (gc GetClient) getClientFromDB(c *gin.Context, repo database.ClientRepository, id int) (cl client.Client, ok bool) {
	cl, err := repo.GetOne(id)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
