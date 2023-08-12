package handlers

import (
	// ... (import statements)

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

func TestSearch_Do(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Create a mock product with sample data
		mockProducts := []*product.Product{
			{
				ID:            3,
				ClientID:      1,
				GuideNumber:   newString("ASD234ASD5"),
				VehiclePlate:  newString("ABC-123"),
				Port:          newInt(123),
				Vault:         newInt(123),
				Quantity:      newInt(123),
				ShippingPrice: newFloat64(123.12),
			},
			{
				ID:            3,
				ClientID:      1,
				GuideNumber:   newString("ABC1234567"),
				VehiclePlate:  newString("ABC-123"),
				Port:          newInt(123),
				Vault:         newInt(123),
				Quantity:      newInt(123),
				ShippingPrice: newFloat64(123.12),
			},
		}

		mockRepo := new(MockProductRepository)
		mockRepo.On("Search", mock.Anything).Return([]*product.Product{mockProducts[0]}, nil)

		// Create a mock context with a request parameter
		req, _ := http.NewRequest("GET", "/path?guideNumber=ASD234ASD5", nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				database.PRODUCT_REPOSITORY: mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})
		// Set up the handler and execute the action
		gc := Search{}
		gc.Do(c)

		// Assert the HTTP status code
		assert.Equal(t, http.StatusOK, rec.Code)

		// Decode the response body
		var responseProducts []*product.Product
		err := json.Unmarshal(rec.Body.Bytes(), &responseProducts)
		assert.NoError(t, err)

		// Compare the responseProduct with the mockProduct
		assert.Equal(t, []*product.Product{mockProducts[0]}, responseProducts)
	})
}
