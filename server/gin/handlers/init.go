package handlers

import (
	"github.com/coffemanfp/test/config"
	"github.com/coffemanfp/test/database"
)

var db database.Repositories
var conf config.ConfigInfo

func Init(newDb database.Repositories, newConf config.ConfigInfo) {
	db = newDb
	conf = newConf
}
