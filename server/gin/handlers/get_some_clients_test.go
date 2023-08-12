package handlers

import (
	// ... (import statements)

	"encoding/json"
	"errors"
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

func TestGetSomeClients_Do(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Create a mock client with sample data
		mockClients := []*client.Client{
			{
				ID:      1,
				Name:    "John",
				Surname: "Doe",
				Auth: auth.Auth{
					Username: "johndoe",
				},
			},
			{
				ID:      2,
				Name:    "Jason",
				Surname: "Smith",
				Auth: auth.Auth{
					Username: "jsmith",
				},
			},
		}

		mockRepo := new(MockClientRepository)
		mockRepo.On("Get", mock.Anything).Return(mockClients, nil)

		// Create a mock context with a request parameter
		req, _ := http.NewRequest("GET", "/path", nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				database.CLIENT_REPOSITORY: mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})
		// Set up the handler and execute the action
		gc := GetSomeClients{}
		gc.Do(c)

		// Assert the HTTP status code
		assert.Equal(t, http.StatusOK, rec.Code)

		// Decode the response body
		var responseClients []*client.Client
		err := json.Unmarshal(rec.Body.Bytes(), &responseClients)
		assert.NoError(t, err)

		// Compare the responseClient with the mockClient
		assert.Equal(t, mockClients, responseClients)
	})

	t.Run("NotFound", func(t *testing.T) {
		mockRepo := new(MockClientRepository)
		mockRepo.On("Get", mock.Anything).Return([]*client.Client{}, errors.New("not found"))

		// Create a mock context with a request parameter
		req, _ := http.NewRequest("GET", "/path", nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				database.CLIENT_REPOSITORY: mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})
		// Set up the handler and execute the action
		gc := GetSomeClients{}
		gc.Do(c)

		assert.Empty(t, rec.Body)
		assert.NotEmpty(t, c.Errors)
		assert.Contains(t, c.Errors[0].Error(), "not found")
	})
}
