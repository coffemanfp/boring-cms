package handlers

import (
	"net/http"

	"github.com/coffemanfp/test/auth"
	"github.com/coffemanfp/test/client"
	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/server/errors"
	"github.com/gin-gonic/gin"
)

type Login struct{}

func (l Login) Do(c *gin.Context) {
	client, ok := l.readCredentials(c)
	if !ok {
		return
	}

	// Continue the search login at the database process
	repo, ok := l.getAuthRepository(c)
	if !ok {
		return
	}

	id, hash, ok := l.searchCredentialsInDB(c, client, repo)
	if !ok {
		return
	}

	ok = l.comparePassword(c, hash, client.Auth.Password)
	if !ok {
		return
	}

	token, ok := l.generateToken(c, id)
	if !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (l Login) readCredentials(c *gin.Context) (client client.Client, ok bool) {
	ok = readRequestData(c, &client)
	return
}

func (l Login) getAuthRepository(c *gin.Context) (repo database.AuthRepository, ok bool) {
	return getAuthRepository(c)
}

func (l Login) searchCredentialsInDB(c *gin.Context, client client.Client, repo database.AuthRepository) (id int, hash string, ok bool) {
	id, hash, err := repo.GetIdAndHashedPassword(client.Auth)
	if err != nil {
		err = errors.NewHTTPError(http.StatusUnauthorized, errors.UNAUTHORIZED_ERROR_MESSAGE)
		handleError(c, err)
		return
	}
	ok = true
	return
}

func (l Login) comparePassword(c *gin.Context, hash, password string) (ok bool) {
	err := auth.CompareHashAndPassword(hash, password)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}

func (l Login) generateToken(c *gin.Context, id int) (token string, ok bool) {
	token, err := auth.GenerateToken(id, conf.Server.JWTLifespan, conf.Server.SecretKey)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
