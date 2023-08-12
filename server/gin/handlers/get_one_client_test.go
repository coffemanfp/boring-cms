package handlers

import (
	// ... (import statements)

	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/coffemanfp/docucentertest/auth"
	"github.com/coffemanfp/docucentertest/client"
	"github.com/coffemanfp/docucentertest/config"
	"github.com/coffemanfp/docucentertest/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetClient_Do(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Create a mock client with sample data
		mockClient := client.Client{
			ID:      1,
			Name:    "John",
			Surname: "Doe",
			Auth: auth.Auth{
				Username: "johndoe",
			},
		}

		mockRepo := new(MockClientRepository)
		mockRepo.On("GetOne", mock.Anything).Return(mockClient, nil)

		// Create a mock context with a request parameter
		req, _ := http.NewRequest("GET", "/path/1", nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req
		c.Params = []gin.Param{{Key: "id", Value: strconv.Itoa(mockClient.ID)}}

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				database.CLIENT_REPOSITORY: mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})
		// Set up the handler and execute the action
		gc := GetClient{}
		gc.Do(c)

		// Assert the HTTP status code
		assert.Equal(t, http.StatusOK, rec.Code)

		// Decode the response body
		var responseClient client.Client
		err := json.Unmarshal(rec.Body.Bytes(), &responseClient)
		assert.NoError(t, err)

		// Compare the responseClient with the mockClient
		assert.Equal(t, mockClient, responseClient)
	})

	t.Run("NotFound", func(t *testing.T) {
		mockRepo := new(MockClientRepository)
		mockRepo.On("GetOne", mock.Anything).Return(client.Client{}, errors.New("not found"))

		// Create a mock context with a request parameter
		req, _ := http.NewRequest("GET", "/path/1", nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req
		c.Params = []gin.Param{{Key: "id", Value: "1"}}

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				database.CLIENT_REPOSITORY: mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})
		// Set up the handler and execute the action
		gc := GetClient{}
		gc.Do(c)

		assert.Empty(t, rec.Body)
		assert.NotEmpty(t, c.Errors)
		assert.Contains(t, c.Errors[0].Error(), "not found")
	})
}
