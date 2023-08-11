package psql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/coffemanfp/test/database"
	"github.com/coffemanfp/test/product"
	"github.com/coffemanfp/test/search"
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
			%s(client_id, guide_number, type, joined_at, delivered_at, shipping_price, vehicle_plate, port, vault, quantity)
		values
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		returning
			id
	`, table)
	err = pr.db.QueryRow(query, p.ClientID, p.GuideNumber, p.Type, p.JoinedAt, p.DeliveredAt, p.ShippingPrice, p.VehiclePlate, p.Port, p.Vault, p.Quantity).Scan(&id)
	if err != nil {
		err = errorInRow(table, "insert", err)
	}
	return
}

func (pr ProductRepository) GetOne(id int) (p product.Product, err error) {
	table := "product"
	query := fmt.Sprintf(`
		select
			id, client_id, guide_number, type, joined_at, delivered_at, shipping_price, vehicle_plate, port, vault, quantity
		from
			%s
		where
			id = $1
	`, table)

	err = pr.db.QueryRow(query, id).Scan(&p.ID, &p.ClientID, &p.GuideNumber, &p.Type, &p.JoinedAt, &p.DeliveredAt, &p.ShippingPrice, &p.VehiclePlate, &p.Port, &p.Vault, &p.Quantity)
	if err != nil {
		p = product.Product{}
		err = errorInRow(table, "get", err)
	}
	return
}

func (pr ProductRepository) Get(page int) (ps []*product.Product, err error) {
	table := "product"
	query := fmt.Sprintf(`
		select
			id, client_id, guide_number, type, joined_at, delivered_at, shipping_price, vehicle_plate, port, vault, quantity
		from
			%s
		limit
			$1
		offset
			$2
	`, table)

	limit, offset := parsePagination(page)

	rows, err := pr.db.Query(query, limit, offset)
	if err != nil {
		err = errorInRow(table, "get", err)
		return
	}

	ps = make([]*product.Product, 0)
	for rows.Next() {
		p := new(product.Product)
		err = rows.Scan(&p.ID, &p.ClientID, &p.GuideNumber, &p.Type, &p.JoinedAt, &p.DeliveredAt, &p.ShippingPrice, &p.VehiclePlate, &p.Port, &p.Vault, &p.Quantity)
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

func (pr ProductRepository) Search(srch search.Search) (ps []*product.Product, err error) {
	table := "product"
	query := fmt.Sprintf(`
		select
			id, client_id, guide_number, type, joined_at, delivered_at, shipping_price, vehicle_plate, port, vault, quantity
		from
			%s
		where
			(nullif($1, '') is null or guide_number = $1) and
			(nullif($2, '') is null or type = $2) and
			(nullif($3, '') is null or vehicle_plate = $3) and
			(nullif($4, 0) is null or port = $4) and
			(nullif($5, 0) is null or vault = $5) and

			((nullif($6, 0.00) is null and nullif($7, 0.00) is null) or ($6 <= shipping_price and $7 >= shipping_price)) and
			(nullif($6, 0.00) is null or $6 <= shipping_price) and
			(nullif($7, 0.00) is null or $7 >= shipping_price) and

			(($8::timestamp is null and $9::timestamp is null) or ($8::timestamp <= joined_at and $9::timestamp >= joined_at)) and
			($8::timestamp is null or $8::timestamp <= joined_at) and
			($9::timestamp is null or $9::timestamp >= joined_at) and

			(($10::timestamp is null and $11::timestamp is null) or ($10::timestamp <= delivered_at and $11::timestamp >= delivered_at)) and
			($10::timestamp is null or $10::timestamp <= delivered_at) and
			($11::timestamp is null or $11::timestamp >= delivered_at)
	`, table)

	rows, err := pr.db.Query(query, srch.GuideNumber, srch.Type, srch.VehiclePlate, srch.Port, srch.Vault, srch.PriceRange.Start, srch.PriceRange.End,
		sql.NullTime{
			Time:  srch.JoinedAtRange.Start,
			Valid: srch.JoinedAtRange.Start != time.Time{},
		},
		sql.NullTime{
			Time:  srch.JoinedAtRange.End,
			Valid: srch.JoinedAtRange.End != time.Time{},
		},
		sql.NullTime{
			Time:  srch.DeliveredAtRange.Start,
			Valid: srch.DeliveredAtRange.Start != time.Time{},
		},
		sql.NullTime{
			Time:  srch.DeliveredAtRange.End,
			Valid: srch.DeliveredAtRange.End != time.Time{},
		},
	)
	if err != nil {
		err = errorInRow(table, "get", err)
		return
	}
	ps = make([]*product.Product, 0)
	for rows.Next() {
		p := new(product.Product)
		err = rows.Scan(&p.ID, &p.ClientID, &p.GuideNumber, &p.Type, &p.JoinedAt, &p.DeliveredAt, &p.ShippingPrice, &p.VehiclePlate, &p.Port, &p.Vault, &p.Quantity)
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
