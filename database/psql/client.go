package psql

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/test/client"
	"github.com/coffemanfp/test/database"
)

type ClientRepository struct {
	db *sql.DB
}

func NewClientRepository(conn *PostgreSQLConnector) (repo database.ClientRepository, err error) {
	db, err := conn.getConn()
	if err != nil {
		return
	}
	repo = ClientRepository{
		db: db,
	}
	return
}

func (cr ClientRepository) GetOne(id int) (c client.Client, err error) {
	table := "client"
	query := fmt.Sprintf(`
		select
			id, title, description, list_id, reminder, due_date, repeat, is_done, is_added_to_my_day, is_important, created_at, created_by
		from
			%s
		where
			id = $1
	`, table)

	err = cr.db.QueryRow(query, id).Scan(&c.ID, &c.Name, &c.Surname, &c.CreatedAt, &c.Auth.Username)
	if err != nil {
		c = client.Client{}
		err = errorInRow(table, "get", err)
	}
	return
}

func (cr ClientRepository) Get(page int) (cs []*client.Client, err error) {
	table := "client"
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

	limit, offset := parsePagination(page)

	rows, err := cr.db.Query(query, limit, offset)
	if err != nil {
		err = errorInRow(table, "get", err)
		return
	}

	cs = make([]*client.Client, 0)
	for rows.Next() {
		c := new(client.Client)
		err = rows.Scan(&c.ID, &c.Name, &c.Surname, &c.CreatedAt, &c.Auth.Username)
		if err != nil {
			err = errorInRow(table, "scan", err)
			cs = nil
			return
		}

		cs = append(cs, c)
	}
	err = rows.Err()
	if err != nil {
		cs = nil
		err = errorInRows(table, "scanning", err)
	}
	return
}
