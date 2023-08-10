package database

import "github.com/coffemanfp/test/client"

const CLIENT_REPOSITORY RepositoryID = "CLIENT_REPOSITORYY"

type ClientRepository interface {
	Get(page int) (clients []*client.Client, err error)
	GetOne(id int) (client client.Client, err error)
}
