package gin

import (
	"github.com/coffemanfp/test/config"
	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/server"
	"github.com/coffemanfp/test/server/gin/handlers"
	"github.com/gin-gonic/gin"
)

// GinEngine is a struct that represents the Gin-based HTTP server engine.
type GinEngine struct {
	conf config.ConfigInfo
	db   database.Database
	r    *gin.Engine
}

// New creates a new instance of the GinEngine.
func New(conf config.ConfigInfo, db database.Database) server.Engine {
	// Initialize a new GinEngine instance with the provided configuration and database
	ge := GinEngine{
		conf: conf,
		db:   db,
		r:    gin.New(),
	}

	// Initialize the handlers with the database repositories and configuration
	handlers.Init(ge.db.Repositories, ge.conf)

	// Use CORS middleware to handle cross-origin requests
	ge.r.Use(newCors(ge.conf))
	// Use custom error handling middleware
	ge.r.Use(errorHandler())

	// Create a new route group for version 1 of the API
	v1 := ge.r.Group("/v1")

	// Set up common middlewares for all routes
	ge.setCommonMiddlewares(v1)
	// Set up authentication-related handlers
	ge.setAuthHandlers(v1)
	// Set up product-related handlers
	ge.setProductHandlers(v1)
	// Set up search-related handlers
	ge.setSearchHandlers(v1)
	// Set up client-related handlers
	ge.setClientHandlers(v1)

	// Return the configured Gin engine
	return ge.r
}

// setAuthHandlers configures authentication-related routes and handlers.
func (ge GinEngine) setAuthHandlers(r *gin.RouterGroup) {
	// Create a sub-group for authentication routes
	auth := r.Group("/auth")
	// Configure the login and register endpoints with their respective handlers
	auth.POST("/login", handlers.Login{}.Do)
	auth.POST("/register", handlers.Register{}.Do)
}

// setProductHandlers configures product-related routes and handlers.
func (ge GinEngine) setProductHandlers(r *gin.RouterGroup) {
	// Create a sub-group for product routes
	product := r.Group("/products")
	// Use authorization middleware to protect these routes
	product.Use(authorize(ge.conf.Server.SecretKey))
	// Configure endpoints for getting, creating, updating, and deleting products
	product.GET("/:id", handlers.GetProduct{}.Do)
	product.GET("", handlers.GetSomeProducts{}.Do)
	product.POST("", handlers.CreateProduct{}.Do)
	product.PUT("/:id", handlers.UpdateProduct{}.Do)
	product.DELETE("/:id", handlers.DeleteProduct{}.Do)
}

// setSearchHandlers configures search-related routes and handlers.
func (ge GinEngine) setSearchHandlers(r *gin.RouterGroup) {
	// Create a sub-group for search routes
	product := r.Group("/search")
	// Use authorization middleware to protect this route
	product.Use(authorize(ge.conf.Server.SecretKey))
	// Configure endpoint for searching products
	product.GET("", handlers.Search{}.Do)
}

// setClientHandlers configures client-related routes and handlers.
func (ge GinEngine) setClientHandlers(r *gin.RouterGroup) {
	// Create a sub-group for client routes
	client := r.Group("/clients")
	// Use authorization middleware to protect these routes
	client.Use(authorize(ge.conf.Server.SecretKey))
	// Configure endpoints for getting clients and getting a specific client
	client.GET("", handlers.GetSomeClients{}.Do)
	client.GET("/:id", handlers.GetClient{}.Do)
}

// setCommonMiddlewares configures common middlewares for all routes.
func (ge GinEngine) setCommonMiddlewares(r *gin.RouterGroup) {
	// Use Gin's recovery middleware for handling panics
	r.Use(gin.Recovery())
	// Use custom logger middleware
	r.Use(logger())
}
