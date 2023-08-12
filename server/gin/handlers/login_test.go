package handlers

import (
	// ... (import statements)

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coffemanfp/docucentertest/auth"
	"github.com/coffemanfp/docucentertest/client"
	"github.com/coffemanfp/docucentertest/config"
	"github.com/coffemanfp/docucentertest/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin_Do(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ar := auth.Auth{
			Username: "john_doe",
			Password: "password",
		}
		arJSON, _ := json.Marshal(ar)

		// Create a mock client credentials
		mockClientCredentials := client.Client{
			Auth: ar,
		}

		// Create a mock authentication repository
		mockRepo := new(MockAuthRepository)
		mockRepo.On("GetIdAndHashedPassword", mockClientCredentials.Auth).Return(1, "$2a$04$ELiP4j1x5NW2nSUEIyJWYui1NCZEpjCZ4ZOpods19haBxP.uPXA8y", nil)

		// Set up the handler and execute the action
		login := Login{}

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				database.AUTH_REPOSITORY: mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})
		r := gin.New()
		r.POST("/login", login.Do)

		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(arJSON))
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert the HTTP status code
		assert.Equal(t, http.StatusOK, rec.Code)

		// Decode the response body
		var responseBody map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		// Assert the generated token is present in the response
		assert.NotNil(t, responseBody["token"])
	})

	t.Run("InvalidCredentials", func(t *testing.T) {
		ar := auth.Auth{
			Username: "john_doe",
			Password: "invalid_password",
		}
		arJSON, _ := json.Marshal(ar)

		// Create a mock client credentials
		mockClientCredentials := client.Client{
			Auth: ar,
		}

		// Create a mock authentication repository
		mockRepo := new(MockAuthRepository)
		mockRepo.On("GetIdAndHashedPassword", mockClientCredentials.Auth).Return(1, "$2a$04$ELiP4j1x5NW2nSUEIyJWYui1NCZEpjCZ4ZOpods19haBxP.uPXA8y", nil)

		// Set up the handler and execute the action
		login := Login{}

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				database.AUTH_REPOSITORY: mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})
		r := gin.New()
		r.POST("/login", login.Do)

		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(arJSON))
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert the HTTP status code
		assert.Empty(t, rec.Body)
	})
}
