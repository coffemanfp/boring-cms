package gin

import (
	"github.com/coffemanfp/test/config"
	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/server"
	"github.com/coffemanfp/test/server/gin/handlers"
	"github.com/gin-gonic/gin"
)

type GinEngine struct {
	conf config.ConfigInfo
	db   database.Database
	r    *gin.Engine
}

func New(conf config.ConfigInfo, db database.Database) server.Engine {
	ge := GinEngine{
		conf: conf,
		db:   db,
		r:    gin.New(),
	}

	handlers.Init(ge.db.Repositories, ge.conf)

	ge.r.Use(newCors(ge.conf))
	ge.r.Use(errorHandler())
	v1 := ge.r.Group("/v1")

	ge.setCommonMiddlewares(v1)
	ge.setAuthHandlers(v1)
	ge.setProductHandlers(v1)
	ge.setSearchHandlers(v1)
	ge.setClientHandlers(v1)
	return ge.r
}

func (ge GinEngine) setAuthHandlers(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	auth.POST("/login", handlers.Login{}.Do)
	auth.POST("/register", handlers.Register{}.Do)
}

func (ge GinEngine) setProductHandlers(r *gin.RouterGroup) {
	product := r.Group("/products")
	product.Use(authorize(ge.conf.Server.SecretKey))
	product.GET("/:id", handlers.GetProduct{}.Do)
	product.GET("", handlers.GetSomeProducts{}.Do)
	product.POST("", handlers.CreateProduct{}.Do)
	product.PUT("/:id", handlers.UpdateProduct{}.Do)
	product.DELETE("/:id", handlers.DeleteProduct{}.Do)
}

func (ge GinEngine) setSearchHandlers(r *gin.RouterGroup) {
	product := r.Group("/search")
	product.Use(authorize(ge.conf.Server.SecretKey))
	product.GET("", handlers.Search{}.Do)
}

func (ge GinEngine) setClientHandlers(r *gin.RouterGroup) {
	client := r.Group("/clients")
	client.Use(authorize(ge.conf.Server.SecretKey))
	client.GET("", handlers.GetSomeClients{}.Do)
	client.GET("/:id", handlers.GetClient{}.Do)
}

func (ge GinEngine) setCommonMiddlewares(r *gin.RouterGroup) {
	r.Use(gin.Recovery())
	r.Use(logger())
}
