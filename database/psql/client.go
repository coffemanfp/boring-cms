package psql

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/docucentertest/client"
	"github.com/coffemanfp/docucentertest/database"
)

// ClientRepository is a struct representing a repository for client-related database operations.
type ClientRepository struct {
	db *sql.DB
}

// NewClientRepository creates a new ClientRepository instance.
func NewClientRepository(conn *PostgreSQLConnector) (repo database.ClientRepository, err error) {
	// Get a database connection from the PostgreSQLConnector.
	db, err := conn.getConn()
	if err != nil {
		return
	}
	// Initialize and return the ClientRepository.
	repo = ClientRepository{
		db: db,
	}
	return
}

// GetOne retrieves a single client from the database based on the provided ID.
func (cr ClientRepository) GetOne(id int) (c client.Client, err error) {
	table := "client"
	// SQL query to select client details based on ID.
	query := fmt.Sprintf(`
		select
			id, title, description, list_id, reminder, due_date, repeat, is_done, is_added_to_my_day, is_important, created_at, created_by
		from
			%s
		where
			id = $1
	`, table)

	// Query the database for the client details based on the provided ID.
	err = cr.db.QueryRow(query, id).Scan(&c.ID, &c.Name, &c.Surname, &c.CreatedAt, &c.Auth.Username)
	if err != nil {
		// In case of an error, create an empty client and generate a detailed error message.
		c = client.Client{}
		err = errorInRow(table, "get", err)
	}
	return
}

// Get retrieves a list of clients from the database based on the provided page number.
func (cr ClientRepository) Get(page int) (cs []*client.Client, err error) {
	table := "client"
	// SQL query to select a list of client details with pagination.
	query := fmt.Sprintf(`
		select
			id, name, surname, created_at, username
		from
			%s
		limit
			$1
		offset
			$2
	`, table)

	// Parse pagination parameters from the provided page number.
	limit, offset := parsePagination(page)

	// Query the database for a list of clients with pagination.
	rows, err := cr.db.Query(query, limit, offset)
	if err != nil {
		// In case of an error, generate a detailed error message.
		err = errorInRow(table, "get", err)
		return
	}

	cs = make([]*client.Client, 0)
	for rows.Next() {
		c := new(client.Client)
		// Scan the row's data into the client structure.
		err = rows.Scan(&c.ID, &c.Name, &c.Surname, &c.CreatedAt, &c.Auth.Username)
		if err != nil {
			// In case of an error during scanning, set the list to nil and return the error.
			err = errorInRow(table, "scan", err)
			cs = nil
			return
		}

		// Append the scanned client to the list.
		cs = append(cs, c)
	}
	err = rows.Err()
	if err != nil {
		// In case of an error during rows iteration, set the list to nil and return the error.
		cs = nil
		err = errorInRows(table, "scanning", err)
	}
	return
}
