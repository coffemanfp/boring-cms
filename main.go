package main

import (
	"fmt"
	"log"

	"github.com/coffemanfp/test/config"
	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/database/psql"
	"github.com/coffemanfp/test/server/gin"
)

func main() {
	// Load configuration from environment variables.
	conf, err := config.NewEnvManagerConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Set up the database connection.
	db, err := setUpDatabase(conf)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new server engine using the loaded configuration and database.
	serverEngine := gin.New(conf, db)

	// Start the server on the specified port.
	serverEngine.Run(fmt.Sprintf(":%d", conf.Server.Port))
}

func setUpDatabase(conf config.ConfigInfo) (db database.Database, err error) {
	// Create a new PostgreSQL database connector.
	db.Conn = psql.NewPostgreSQLConnector(
		conf.PostgreSQLProperties.User,
		conf.PostgreSQLProperties.Password,
		conf.PostgreSQLProperties.Name,
		conf.PostgreSQLProperties.Host,
		conf.PostgreSQLProperties.Port,
	)

	// Connect to the database.
	err = db.Conn.Connect()
	if err != nil {
		return
	}

	// Create a new authentication repository using the PostgreSQL connector.
	authRepo, err := psql.NewAuthRepository(db.Conn.(*psql.PostgreSQLConnector))
	if err != nil {
		return
	}

	// Create a new client repository using the PostgreSQL connector.
	clientRepo, err := psql.NewClientRepository(db.Conn.(*psql.PostgreSQLConnector))
	if err != nil {
		return
	}

	// Create a new product repository using the PostgreSQL connector.
	productRepo, err := psql.NewProductRepository(db.Conn.(*psql.PostgreSQLConnector))
	if err != nil {
		return
	}

	// Initialize the database repositories.
	db.Repositories = map[database.RepositoryID]interface{}{
		database.AUTH_REPOSITORY:    authRepo,
		database.CLIENT_REPOSITORY:  clientRepo,
		database.PRODUCT_REPOSITORY: productRepo,
	}
	return
}
