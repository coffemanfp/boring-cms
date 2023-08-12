package psql

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/test/auth"
	"github.com/coffemanfp/test/client"
	"github.com/coffemanfp/test/database"
)

// AuthRepository is a struct representing a repository for authentication-related database operations.
type AuthRepository struct {
	db *sql.DB
}

// NewAuthRepository creates a new AuthRepository instance.
func NewAuthRepository(conn *PostgreSQLConnector) (repo database.AuthRepository, err error) {
	// Get a database connection from the connector.
	db, err := conn.getConn()
	if err != nil {
		return
	}
	// Initialize and return the AuthRepository.
	repo = AuthRepository{
		db: db,
	}
	return
}

// GetIdAndHashedPassword retrieves the client's ID and hashed password from the database based on the provided auth credentials.
func (ar AuthRepository) GetIdAndHashedPassword(auth auth.Auth) (id int, hashed string, err error) {
	table := "client"
	query := `
		select id, password from client where username = $1
	`

	// Query the database for the ID and hashed password based on the provided username.
	err = ar.db.QueryRow(query, auth.Username).Scan(&id, &hashed)
	if err != nil {
		err = errorInRow(table, "get", err)
	}
	return
}

// Register registers a new client in the database and returns the assigned ID.
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

	// Insert the new client's details into the database and retrieve the assigned ID.
	err = ar.db.QueryRow(query, client.Name, client.Surname, client.Auth.Username, client.Auth.Password, client.CreatedAt).Scan(&id)
	if err != nil {
		err = errorInRow(table, "insert", err)
	}
	return
}
