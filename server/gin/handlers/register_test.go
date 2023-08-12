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
	"github.com/stretchr/testify/mock"
)

func TestRegister_Do(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ar := auth.Auth{
			Username: "john_doe",
			Password: "password",
		}

		// Create a mock client credentials
		mockClientCredentials := client.Client{
			Auth:    ar,
			Name:    "John",
			Surname: "Smith",
		}
		crJSON, _ := json.Marshal(mockClientCredentials)

		// Create a mock authentication repository
		mockRepo := new(MockAuthRepository)
		mockRepo.On("Register", mock.Anything).Return(1, nil)

		// Set up the handler and execute the action
		register := Register{}

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				database.AUTH_REPOSITORY: mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})
		r := gin.New()
		r.POST("/register", register.Do)

		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(crJSON))
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert the HTTP status code
		assert.Equal(t, http.StatusCreated, rec.Code)

		// Decode the response body
		var responseBody map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		// Assert the generated token is present in the response
		assert.NotNil(t, responseBody["token"])
	})

	t.Run("InvalidData", func(t *testing.T) {
		ar := auth.Auth{
			Username: "j@hn_doe",
			Password: "password",
		}

		// Create a mock client credentials
		mockClientCredentials := client.Client{
			Auth:    ar,
			Name:    "John",
			Surname: "Smith",
		}
		crJSON, _ := json.Marshal(mockClientCredentials)

		// Create a mock authentication repository
		mockRepo := new(MockAuthRepository)

		// Set up the handler and execute the action
		register := Register{}

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				database.AUTH_REPOSITORY: mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})
		r := gin.New()
		r.POST("/register", register.Do)

		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(crJSON))
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		assert.Empty(t, rec.Body)
	})
}
