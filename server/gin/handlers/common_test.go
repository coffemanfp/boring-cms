package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coffemanfp/test/auth"
	"github.com/coffemanfp/test/client"
	"github.com/coffemanfp/test/config"
	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/product"
	"github.com/coffemanfp/test/search"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestReadRequestData(t *testing.T) {
	t.Run("SuccessfulBinding", func(t *testing.T) {
		// Create a sample JSON payload
		jsonData := `{"id": 1, "name": "Product A", "price": 10.5}`
		req, _ := http.NewRequest("POST", "/", bytes.NewBufferString(jsonData))
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		// Create a struct to bind the JSON data
		var product struct {
			ID    int     `json:"id"`
			Name  string  `json:"name"`
			Price float64 `json:"price"`
		}

		// Try to read and bind JSON data
		ok := readRequestData(c, &product)

		assert.True(t, ok)
		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, 1, product.ID)
		assert.Equal(t, "Product A", product.Name)
		assert.Equal(t, 10.5, product.Price)
	})

	t.Run("FailedBinding", func(t *testing.T) {
		// Create an invalid JSON payload
		jsonData := `{"id "invalid", ame": "Product A", "price": 10.`
		req, _ := http.NewRequest("POST", "/", bytes.NewBufferString(jsonData))
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		// Create a struct to bind the JSON data
		var product struct {
			ID    int     `json:"id"`
			Name  string  `json:"name"`
			Price float64 `json:"price"`
		}

		// Try to read and bind JSON data
		ok := readRequestData(c, &product)

		assert.False(t, ok)
		assert.False(t, ok)
		assert.NotEmpty(t, c.Errors)
		assert.Contains(t, c.Errors[0].Error(), "invalid character")
	})
}

type MockAuthRepository struct {
	mock.Mock
}

func (m *MockAuthRepository) GetIdAndHashedPassword(auth auth.Auth) (int, string, error) {
	args := m.Called(auth)
	return args.Int(0), args.String(1), args.Error(2)
}

func (m *MockAuthRepository) Register(client client.Client) (int, error) {
	args := m.Called(client)
	return args.Int(0), args.Error(1)
}

func TestGetAuthRepository(t *testing.T) {
	t.Run("SuccessfulRepositoryRetrieval", func(t *testing.T) {
		mockRepo := new(MockAuthRepository)

		req, _ := http.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				database.AUTH_REPOSITORY: mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})

		repo, ok := getAuthRepository(c)

		assert.True(t, ok)
		assert.NotNil(t, repo)
		assert.NoError(t, rec.Result().Body.Close())
	})

	t.Run("FailedRepositoryRetrieval", func(t *testing.T) {
		mockRepo := new(MockAuthRepository)

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				"Invalid Repository ID": mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})
		req, _ := http.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		repo, ok := getAuthRepository(c)

		assert.False(t, ok)
		assert.Nil(t, repo)
		assert.NoError(t, rec.Result().Body.Close())
	})
}

type MockClientRepository struct {
	mock.Mock
}

func (m *MockClientRepository) GetOne(id int) (client.Client, error) {
	args := m.Called(id)
	return args.Get(0).(client.Client), args.Error(1)
}

func (m *MockClientRepository) Get(page int) ([]*client.Client, error) {
	args := m.Called(page)
	return args.Get(0).([]*client.Client), args.Error(1)
}

func TestGetClientRepository(t *testing.T) {
	t.Run("SuccessfulRepositoryRetrieval", func(t *testing.T) {
		mockRepo := new(MockClientRepository)

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				database.CLIENT_REPOSITORY: mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})

		req, _ := http.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		repo, ok := getClientRepository(c)

		assert.True(t, ok)
		assert.NotNil(t, repo)
		assert.NoError(t, rec.Result().Body.Close())
	})

	t.Run("FailedRepositoryRetrieval", func(t *testing.T) {
		mockRepo := new(MockClientRepository)

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				"Invalid Repository ID": mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})

		req, _ := http.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		repo, ok := getClientRepository(c)

		assert.False(t, ok)
		assert.Nil(t, repo)
		assert.NoError(t, rec.Result().Body.Close())
	})
}

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Get(page, clientID int) ([]*product.Product, error) {
	args := m.Called(page, clientID)
	return args.Get(0).([]*product.Product), args.Error(1)
}

func (m *MockProductRepository) GetOne(id, clientID int) (product.Product, error) {
	args := m.Called(id, clientID)
	return args.Get(0).(product.Product), args.Error(1)
}

func (m *MockProductRepository) Create(product product.Product) (int, error) {
	args := m.Called(product)
	return args.Int(0), args.Error(1)
}

func (m *MockProductRepository) Search(search search.Search) ([]*product.Product, error) {
	args := m.Called(search)
	return args.Get(0).([]*product.Product), args.Error(1)
}

func (m *MockProductRepository) Update(product product.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(id, clientID int) error {
	args := m.Called(id, clientID)
	return args.Error(0)
}

func TestGetProductRepository(t *testing.T) {
	t.Run("SuccessfulRepositoryRetrieval", func(t *testing.T) {
		mockRepo := new(MockProductRepository)

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				database.PRODUCT_REPOSITORY: mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})
		req, _ := http.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		repo, ok := getProductRepository(c)

		assert.True(t, ok)
		assert.NotNil(t, repo)
		assert.NoError(t, rec.Result().Body.Close())
	})

	t.Run("FailedRepositoryRetrieval", func(t *testing.T) {
		mockRepo := new(MockProductRepository)

		db := database.Database{
			Repositories: map[database.RepositoryID]interface{}{
				"Invalid Repository Key": mockRepo,
			},
		}

		Init(db.Repositories, config.ConfigInfo{})

		req, _ := http.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		repo, ok := getProductRepository(c)

		assert.False(t, ok)
		assert.Nil(t, repo)
		assert.NoError(t, rec.Result().Body.Close())
	})
}

func TestReadIntFromURL(t *testing.T) {
	t.Run("EmptyParameter", func(t *testing.T) {
		r := gin.New()
		r.GET("/path/:id", func(c *gin.Context) {
			c.String(200, "pong")
		})

		req, _ := http.NewRequest("GET", "/path/0", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		c := gin.CreateTestContextOnly(rec, r)
		c.Params = append(c.Params, gin.Param{
			Key:   "id",
			Value: "0",
		})
		v, ok := readIntFromURL(c, "id", false)
		assert.True(t, ok)
		assert.Equal(t, 200, rec.Code)
		assert.Empty(t, v)
		assert.NoError(t, rec.Result().Body.Close())
	})

	t.Run("QueryParam", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/path?id=456", nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		v, ok := readIntFromURL(c, "id", true)

		assert.True(t, ok)
		assert.Equal(t, 456, v)
		assert.NoError(t, rec.Result().Body.Close())
	})

	t.Run("PathParam", func(t *testing.T) {
		r := gin.New()
		r.GET("/path/:id", func(c *gin.Context) {
			c.String(200, "pong")
		})

		req, _ := http.NewRequest("GET", "/path/123", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		c := gin.CreateTestContextOnly(rec, r)
		c.Params = append(c.Params, gin.Param{
			Key:   "id",
			Value: "123",
		})
		v, ok := readIntFromURL(c, "id", false)
		assert.True(t, ok)
		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, 123, v)
		assert.NoError(t, rec.Result().Body.Close())
	})

	t.Run("InvalidParameter", func(t *testing.T) {
		r := gin.New()
		r.GET("/path/:id", func(c *gin.Context) {
			c.String(200, "pong")
		})

		req, _ := http.NewRequest("GET", "/path/abc", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		c := gin.CreateTestContextOnly(rec, r)
		c.Params = append(c.Params, gin.Param{
			Key:   "id",
			Value: "abc",
		})
		v, ok := readIntFromURL(c, "id", false)
		assert.False(t, ok)
		assert.Equal(t, 200, rec.Code)
		assert.Empty(t, v)
		assert.NoError(t, rec.Result().Body.Close())
	})
}

func TestReadFloat64FromURL(t *testing.T) {
	t.Run("EmptyParameter", func(t *testing.T) {
		r := gin.New()
		r.GET("/path/:id", func(c *gin.Context) {
			c.String(200, "pong")
		})

		req, _ := http.NewRequest("GET", "/path/0", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		c := gin.CreateTestContextOnly(rec, r)
		c.Params = append(c.Params, gin.Param{
			Key:   "id",
			Value: "0",
		})
		v, ok := readFloatFromURL(c, "id", false)
		assert.True(t, ok)
		assert.Equal(t, 200, rec.Code)
		assert.Empty(t, v)
		assert.NoError(t, rec.Result().Body.Close())
	})

	t.Run("QueryParam", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/path?id=456.56", nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		v, ok := readFloatFromURL(c, "id", true)

		assert.True(t, ok)
		assert.Equal(t, 456.56, v)
		assert.NoError(t, rec.Result().Body.Close())
	})

	t.Run("PathParam", func(t *testing.T) {
		r := gin.New()
		r.GET("/path/:id", func(c *gin.Context) {
			c.String(200, "pong")
		})

		req, _ := http.NewRequest("GET", "/path/123.56", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		c := gin.CreateTestContextOnly(rec, r)
		c.Params = append(c.Params, gin.Param{
			Key:   "id",
			Value: "123.56",
		})
		v, ok := readFloatFromURL(c, "id", false)
		assert.True(t, ok)
		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, 123.56, v)
		assert.NoError(t, rec.Result().Body.Close())
	})

	t.Run("InvalidParameter", func(t *testing.T) {
		r := gin.New()
		r.GET("/path/:id", func(c *gin.Context) {
			c.String(200, "pong")
		})

		req, _ := http.NewRequest("GET", "/path/abc", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		c := gin.CreateTestContextOnly(rec, r)
		c.Params = append(c.Params, gin.Param{
			Key:   "id",
			Value: "abc",
		})
		v, ok := readFloatFromURL(c, "id", false)
		assert.False(t, ok)
		assert.Equal(t, 200, rec.Code)
		assert.Empty(t, v)
		assert.NoError(t, rec.Result().Body.Close())
	})
}

func newString(s string) *string {
	n := &s
	return n
}

func newInt(i int) *int {
	n := &i
	return n
}

func newFloat64(f float64) *float64 {
	n := &f
	return n
}
