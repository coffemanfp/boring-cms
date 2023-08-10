package handlers

import (
	"net/http"

	"github.com/coffemanfp/test/auth"
	"github.com/coffemanfp/test/client"
	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/server/errors"
	"github.com/gin-gonic/gin"
)

type Register struct{}

func (r Register) Do(c *gin.Context) {
	client, ok := r.readClient(c)
	if !ok {
		return
	}

	client, ok = r.createNewClient(c, client)
	if !ok {
		return
	}

	repo, ok := r.getAuthRepository(c)
	if !ok {
		return
	}

	id, ok := r.registerClientInDB(c, client, repo)
	if !ok {
		return
	}

	token, ok := r.generateToken(c, id)
	if !ok {
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})

}

func (r Register) readClient(c *gin.Context) (client client.Client, ok bool) {
	ok = readRequestData(c, &client)
	return
}

func (r Register) createNewClient(c *gin.Context, clientR client.Client) (cl client.Client, ok bool) {
	cl, err := client.New(clientR)
	if err != nil {
		err = errors.NewHTTPError(http.StatusBadRequest, err.Error())
		handleError(c, err)
		return
	}
	ok = true
	return
}

func (r Register) getAuthRepository(c *gin.Context) (repo database.AuthRepository, ok bool) {
	return getAuthRepository(c)
}

func (r Register) registerClientInDB(c *gin.Context, client client.Client, repo database.AuthRepository) (id int, ok bool) {
	id, err := repo.Register(client)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}

func (r Register) generateToken(c *gin.Context, id int) (token string, ok bool) {
	token, err := auth.GenerateToken(id, conf.Server.JWTLifespan, conf.Server.SecretKey)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
