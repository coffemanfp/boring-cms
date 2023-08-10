package psql

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/product"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(conn *PostgreSQLConnector) (repo database.ProductRepository, err error) {
	db, err := conn.getConn()
	if err != nil {
		return
	}
	repo = ProductRepository{
		db: db,
	}
	return
}

func (pr ProductRepository) Create(p product.Product) (id int, err error) {
	table := "product"
	query := fmt.Sprintf(`
		insert into
			%s(client_id, guide_number, type, joined_at, delivered_at, shipping_price, vehicle_plate, port, vault)
		values
			($1, $2, $3, $4, $5, $6, $7, $8, $9)
		returning
			id
	`, table)
	err = pr.db.QueryRow(query, p.ClientID, p.GuideNumber, &p.Type, p.JoinedAt, p.DeliveredAt, p.ShippingPrice, p.VehiclePlate, p.Port, p.Vault).Scan(&id)
	if err != nil {
		err = errorInRow(table, "insert", err)
	}
	return
}

func (pr ProductRepository) GetOne(id int) (p product.Product, err error) {
	table := "product"
	query := fmt.Sprintf(`
		select
			id, client_id, guide_number, type, joined_at, delivered_at, shipping_price, vehicle_plate, port, vault
		from
			%s
		where
			id = $1
	`, table)

	err = pr.db.QueryRow(query, id).Scan(&p.ID, p.ClientID, p.GuideNumber, &p.Type, p.JoinedAt, p.DeliveredAt, p.ShippingPrice, p.VehiclePlate, p.Port, p.Vault)
	if err != nil {
		p = product.Product{}
		err = errorInRow(table, "get", err)
	}
	return
}

func (tr ProductRepository) Get(page int) (ps []*product.Product, err error) {
	table := "product"
	query := fmt.Sprintf(`
		select
			id, client_id, guide_number, type, joined_at, delivered_at, shipping_price, vehicle_plate, port, vault
		from
			%s
		limit
			$1
		offset
			$2
	`, table)

	limit, offset := parsePagination(page)

	rows, err := tr.db.Query(query, limit, offset)
	if err != nil {
		err = errorInRow(table, "get", err)
		return
	}

	ps = make([]*product.Product, 0)
	for rows.Next() {
		p := new(product.Product)
		err = rows.Scan(&p.ID, &p.ID, p.ClientID, p.GuideNumber, &p.Type, p.JoinedAt, p.DeliveredAt, p.ShippingPrice, p.VehiclePlate, p.Port, p.Vault)
		if err != nil {
			err = errorInRow(table, "scan", err)
			ps = nil
			return
		}

		ps = append(ps, p)
	}
	err = rows.Err()
	if err != nil {
		ps = nil
		err = errorInRows(table, "scanning", err)
	}
	return
}
