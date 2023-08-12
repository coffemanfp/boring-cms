package handlers

import (
	"github.com/coffemanfp/docucentertest/config"
	"github.com/coffemanfp/docucentertest/database"
)

var db database.Repositories
var conf config.ConfigInfo

// Init initializes the global database and configuration variables.
// It sets the provided repositories and configuration information to be used throughout the application.
func Init(newDb database.Repositories, newConf config.ConfigInfo) {
	db = newDb
	conf = newConf
}
