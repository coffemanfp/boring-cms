package psql

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/test/auth"
	"github.com/coffemanfp/test/client"
	"github.com/coffemanfp/test/database"
)

type AuthRepository struct {
	db *sql.DB
}

// NewAuthRepository initializes a new auth repository instance.
//
//	@param conn *PostgreSQLConnector: is the PostgreSQLConnector handler.
//	@return repo database.AuthRepository: is the final interface to keep
//	 the AuthRepository implementation.
//	@return err error: database connection error.
func NewAuthRepository(conn *PostgreSQLConnector) (repo database.AuthRepository, err error) {
	db, err := conn.getConn()
	if err != nil {
		return
	}
	repo = AuthRepository{
		db: db,
	}
	return
}

func (ar AuthRepository) GetIdAndHashedPassword(auth auth.Auth) (id int, hashed string, err error) {
	table := "client"
	query := `
		select id, password from client where username = $1
	`

	err = ar.db.QueryRow(query, auth.Username).Scan(&id, &hashed)
	if err != nil {
		err = errorInRow(table, "get", err)
	}
	return
}

func (ar AuthRepository) Register(client client.Client) (id int, err error) {
	table := "client"
	query := fmt.Sprintf(`
		insert into
			%s(name, surname, username, password, created_at)
		values
			($1, $2, $3, $4, $5)
		returning
			id
	`, table)

	err = ar.db.QueryRow(query, client.Name, client.Surname, client.Auth.Username, client.Auth.Password, client.CreatedAt).Scan(&id)
	if err != nil {
		err = errorInRow(table, "insert", err)
	}
	return
}
