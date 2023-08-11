package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/server/errors"
	"github.com/gin-gonic/gin"
)

func readRequestData(c *gin.Context, v interface{}) (ok bool) {
	err := c.ShouldBindJSON(v)
	if err != nil {
		err = errors.NewHTTPError(http.StatusBadRequest, err.Error())
		handleError(c, err)
		return
	}
	ok = true
	return
}

func getAuthRepository(c *gin.Context) (repo database.AuthRepository, ok bool) {
	repo, err := database.GetRepository[database.AuthRepository](db, database.AUTH_REPOSITORY)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}

func getClientRepository(c *gin.Context) (repo database.ClientRepository, ok bool) {
	repo, err := database.GetRepository[database.ClientRepository](db, database.CLIENT_REPOSITORY)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}

func getProductRepository(c *gin.Context) (repo database.ProductRepository, ok bool) {
	repo, err := database.GetRepository[database.ProductRepository](db, database.PRODUCT_REPOSITORY)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}

func readIntFromURL(c *gin.Context, param string, isQueryParam bool) (v int, ok bool) {
	var p string
	if isQueryParam {
		p = c.Query(param)
	} else {
		p = c.Param(param)
	}
	if p == "" {
		ok = true
		return
	}
	v, err := strconv.Atoi(p)
	if err != nil {
		err = fmt.Errorf("invalid %s param: %s", param, p)
		err = errors.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		handleError(c, err)
		return
	}

	ok = true
	return
}

func readFloatFromURL(c *gin.Context, param string, isQueryParam bool) (v float64, ok bool) {
	var p string
	if isQueryParam {
		p = c.Query(param)
	} else {
		p = c.Param(param)

	}
	if p == "" {
		ok = true
		return
	}
	v, err := strconv.ParseFloat(p, 64)
	if err != nil {
		err = fmt.Errorf("invalid %s param: %s", param, p)
		err = errors.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		handleError(c, err)
		return
	}

	ok = true
	return
}

func readPagination(c *gin.Context) (page int, ok bool) {
	page, ok = readIntFromURL(c, "page", false)
	return
}

func handleError(c *gin.Context, err error) {
	c.Error(err)
	c.Abort()
}
