package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coffemanfp/docucentertest/config"
	"github.com/coffemanfp/docucentertest/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteProduct_Do(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockProductRepository)
		mockRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		ct := DeleteProduct{}

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				database.PRODUCT_REPOSITORY: mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})
		r := gin.New()
		r.DELETE("/path/:id", ct.Do)

		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/path/1", nil)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Empty(t, rec.Body)
	})

	t.Run("NotFound", func(t *testing.T) {
		mockRepo := new(MockClientRepository)
		mockRepo.On("Delete", mock.Anything, mock.Anything).Return(errors.New("not found"))

		// Create a mock context with a request parameter
		req, _ := http.NewRequest("DELETE", "/path/1", nil)
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
		gc := DeleteProduct{}
		gc.Do(c)

		assert.Empty(t, rec.Body)
		assert.NotEmpty(t, c.Errors)
		assert.Contains(t, c.Errors[0].Error(), "not found")
	})

}
