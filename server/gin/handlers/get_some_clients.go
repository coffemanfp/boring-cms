package handlers

import (
	"net/http"

	"github.com/coffemanfp/test/client"
	"github.com/coffemanfp/test/database"
	"github.com/gin-gonic/gin"
)

type GetSomeClients struct{}

func (gst GetSomeClients) Do(c *gin.Context) {
	page, ok := readPagination(c)
	if !ok {
		return
	}

	repo, ok := getClientRepository(c)
	if !ok {
		return
	}

	cs, ok := gst.get(c, repo, page)
	if !ok {
		return
	}

	c.JSON(http.StatusOK, cs)
}

func (gst GetSomeClients) get(c *gin.Context, repo database.ClientRepository, page int) (cs []*client.Client, ok bool) {
	cs, err := repo.Get(page)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
