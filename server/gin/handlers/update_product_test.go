package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coffemanfp/docucentertest/config"
	"github.com/coffemanfp/docucentertest/database"
	"github.com/coffemanfp/docucentertest/product"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateProduct_Do(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockProductRepository)
		mockRepo.On("Update", mock.Anything).Return(nil)

		pr := product.Product{
			ID:           3,
			ClientID:     1,
			GuideNumber:  newString("ABC1234567"),
			VehiclePlate: newString("ABC-123"),
		}
		prJSON, _ := json.Marshal(pr)
		ct := UpdateProduct{}

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				database.PRODUCT_REPOSITORY: mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})
		r := gin.New()
		r.PUT("/path/:id", ct.Do)

		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/path/3", bytes.NewBuffer(prJSON))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Empty(t, rec.Body)
	})

	t.Run("InvalidData", func(t *testing.T) {
		mockRepo := new(MockProductRepository)
		mockRepo.On("Update", mock.Anything).Return(nil)

		pr := product.Product{
			ID:       3,
			ClientID: 1,
		}
		prJSON, _ := json.Marshal(pr)
		ct := UpdateProduct{}

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				database.PRODUCT_REPOSITORY: mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})
		r := gin.New()
		r.PUT("/path/3", ct.Do)

		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/path/3", bytes.NewBuffer(prJSON))
		r.ServeHTTP(rec, req)

		assert.Empty(t, rec.Body)
	})
}
