package handlers

import (
	// ... (import statements)

	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/coffemanfp/docucentertest/config"
	"github.com/coffemanfp/docucentertest/database"
	"github.com/coffemanfp/docucentertest/product"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetProduct_Do(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Create a mock product with sample data
		mockProduct := product.Product{
			ID:            3,
			ClientID:      1,
			GuideNumber:   newString("ABC1234567"),
			VehiclePlate:  newString("ABC-123"),
			Port:          newInt(123),
			Vault:         newInt(123),
			Quantity:      newInt(123),
			ShippingPrice: newFloat64(123.12),
		}

		mockRepo := new(MockProductRepository)
		mockRepo.On("GetOne", mock.Anything, mock.Anything).Return(mockProduct, nil)

		// Create a mock context with a request parameter
		req, _ := http.NewRequest("GET", "/path/3", nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req
		c.Params = []gin.Param{{Key: "id", Value: strconv.Itoa(mockProduct.ID)}}

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				database.PRODUCT_REPOSITORY: mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})
		// Set up the handler and execute the action
		gc := GetProduct{}
		gc.Do(c)

		// Assert the HTTP status code
		assert.Equal(t, http.StatusOK, rec.Code)

		// Decode the response body
		var responseProduct product.Product
		err := json.Unmarshal(rec.Body.Bytes(), &responseProduct)
		assert.NoError(t, err)

		// Compare the responseProduct with the mockProduct
		if assert.NotEmpty(t, responseProduct.Discount) {
			mockProduct.Discount = responseProduct.Discount
		}
		assert.Equal(t, mockProduct, responseProduct)
	})

	t.Run("NotFound", func(t *testing.T) {
		mockRepo := new(MockProductRepository)
		mockRepo.On("GetOne", mock.Anything, mock.Anything).Return(product.Product{}, errors.New("not found"))

		// Create a mock context with a request parameter
		req, _ := http.NewRequest("GET", "/path/1", nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req
		c.Params = []gin.Param{{Key: "id", Value: "1"}}

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				database.PRODUCT_REPOSITORY: mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})
		// Set up the handler and execute the action
		gc := GetProduct{}
		gc.Do(c)

		assert.Empty(t, rec.Body)
		assert.NotEmpty(t, c.Errors)
		assert.Contains(t, c.Errors[0].Error(), "not found")
	})
}
