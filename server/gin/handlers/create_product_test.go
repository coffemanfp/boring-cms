package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coffemanfp/test/config"
	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/product"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProduct_Do(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		newID := 1
		mockRepo := new(MockProductRepository)
		mockRepo.On("Create", mock.Anything).Return(newID, nil)

		pr := product.Product{
			ClientID:     1,
			GuideNumber:  newString("ABC1234567"),
			VehiclePlate: newString("ABC-123"),
		}
		prJSON, _ := json.Marshal(pr)
		ct := CreateProduct{}

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				database.PRODUCT_REPOSITORY: mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})
		r := gin.New()
		r.POST("/path", ct.Do)

		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/path", bytes.NewBuffer(prJSON))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)

		var responseProduct product.Product
		err := json.Unmarshal(rec.Body.Bytes(), &responseProduct)
		assert.NoError(t, err)
		pr.ID = newID
		assert.Equal(t, pr, responseProduct)
	})

	t.Run("InvalidData", func(t *testing.T) {
		mockRepo := new(MockProductRepository)
		mockRepo.On("Create", mock.Anything).Return(0, nil)

		pr := product.Product{
			ClientID: 1,
		}
		prJSON, _ := json.Marshal(pr)
		ct := CreateProduct{}

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				database.PRODUCT_REPOSITORY: mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})
		r := gin.New()
		r.POST("/path", ct.Do)

		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/path", bytes.NewBuffer(prJSON))
		r.ServeHTTP(rec, req)

		assert.Empty(t, rec.Body)
	})
}
