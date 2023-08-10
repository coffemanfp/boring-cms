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
	var err error
	conf, err := config.NewEnvManagerConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := setUpDatabase(conf)
	if err != nil {
		log.Fatal(err)
	}

	serverEngine := gin.New(conf, db)
	serverEngine.Run(fmt.Sprintf(":%d", conf.Server.Port))
}

func setUpDatabase(conf config.ConfigInfo) (db database.Database, err error) {
	db.Conn = psql.NewPostgreSQLConnector(
		conf.PostgreSQLProperties.User,
		conf.PostgreSQLProperties.Password,
		conf.PostgreSQLProperties.Name,
		conf.PostgreSQLProperties.Host,
		conf.PostgreSQLProperties.Port,
	)

	err = db.Conn.Connect()
	if err != nil {
		return
	}

	authRepo, err := psql.NewAuthRepository(db.Conn.(*psql.PostgreSQLConnector))
	if err != nil {
		return
	}

	clientRepo, err := psql.NewClientRepository(db.Conn.(*psql.PostgreSQLConnector))
	if err != nil {
		return
	}

	db.Repositories = map[database.RepositoryID]interface{}{
		database.AUTH_REPOSITORY:   authRepo,
		database.CLIENT_REPOSITORY: clientRepo,
	}
	return
}
