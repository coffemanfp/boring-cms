package psql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/coffemanfp/docucentertest/database"
	"github.com/coffemanfp/docucentertest/product"
	"github.com/coffemanfp/docucentertest/search"
)

// ProductRepository represents a repository for managing products in PostgreSQL.
type ProductRepository struct {
	db *sql.DB
}

// NewProductRepository creates a new ProductRepository instance using a PostgreSQL connector.
func NewProductRepository(conn *PostgreSQLConnector) (repo database.ProductRepository, err error) {
	// Establish a database connection using the provided connector.
	db, err := conn.getConn()
	if err != nil {
		return
	}
	// Create and return a new ProductRepository with the established connection.
	repo = ProductRepository{
		db: db,
	}
	return
}

// Create inserts a new product into the database and returns its ID.
func (pr ProductRepository) Create(p product.Product) (id int, err error) {
	table := "product"
	// Define the SQL query for inserting a new product.
	query := fmt.Sprintf(`
		insert into
			%s(client_id, guide_number, type, joined_at, delivered_at, shipping_price, vehicle_plate, port, vault, quantity)
		values
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		returning
			id
	`, table)
	// Execute the query and scan the result into the 'id' variable.
	err = pr.db.QueryRow(query, p.ClientID, p.GuideNumber, p.Type, p.JoinedAt, p.DeliveredAt, p.ShippingPrice, p.VehiclePlate, p.Port, p.Vault, p.Quantity).Scan(&id)
	if err != nil {
		// If an error occurs, wrap it with a descriptive error message and code.
		err = errorInRow(table, "insert", err)
	}
	return
}

// GetOne retrieves a single product by its ID and clientID from the database.
func (pr ProductRepository) GetOne(id, clientID int) (p product.Product, err error) {
	table := "product"
	// Define the SQL query for retrieving a product by ID and clientID.
	query := fmt.Sprintf(`
		select
			id, client_id, guide_number, type, joined_at, delivered_at, shipping_price, vehicle_plate, port, vault, quantity
		from
			%s
		where
			id = $1 and client_id = $2
	`, table)

	// Execute the query and scan the result into the 'p' variable.
	err = pr.db.QueryRow(query, id, clientID).Scan(&p.ID, &p.ClientID, &p.GuideNumber, &p.Type, &p.JoinedAt,
		&p.DeliveredAt, &p.ShippingPrice, &p.VehiclePlate, &p.Port, &p.Vault, &p.Quantity)
	if err != nil {
		// If an error occurs, set 'p' to a default product and wrap the error with additional information.
		p = product.Product{}
		err = errorInRow(table, "get", err)
	}
	return
}

// Get retrieves a list of products for a given page and clientID from the database.
func (pr ProductRepository) Get(page, clientID int) (ps []*product.Product, err error) {
	table := "product"
	// Define the SQL query for retrieving products for a specific client, with pagination.
	query := fmt.Sprintf(`
		select
			id, client_id, guide_number, type, joined_at, delivered_at, shipping_price, vehicle_plate, port, vault, quantity
		from
			%s
		where
			client_id = $3
		limit
			$1
		offset
			$2
	`, table)

	// Calculate the 'limit' and 'offset' values based on the page number.
	limit, offset := parsePagination(page)

	// Execute the query and retrieve rows from the database.
	rows, err := pr.db.Query(query, limit, offset, clientID)
	if err != nil {
		// If an error occurs while querying, wrap it with a meaningful error message and code.
		err = errorInRow(table, "get", err)
		return
	}

	// Initialize a slice to store the retrieved products.
	ps = make([]*product.Product, 0)
	// Iterate through each row of the result set.
	for rows.Next() {
		p := new(product.Product)
		// Scan the row's columns into the 'p' variable.
		err = rows.Scan(&p.ID, &p.ClientID, &p.GuideNumber, &p.Type, &p.JoinedAt, &p.DeliveredAt, &p.ShippingPrice, &p.VehiclePlate, &p.Port, &p.Vault, &p.Quantity)
		if err != nil {
			// If an error occurs during scanning, wrap it with additional error information.
			err = errorInRow(table, "scan", err)
			ps = nil
			return
		}

		// Append the scanned product to the 'ps' slice.
		ps = append(ps, p)
	}
	// Check for any error that occurred during iteration.
	err = rows.Err()
	if err != nil {
		// If an error occurred while iterating through rows, wrap it with additional error information.
		ps = nil
		err = errorInRows(table, "scanning", err)
	}
	return
}

// Search searches for products based on the provided search criteria.
func (pr ProductRepository) Search(srch search.Search) (ps []*product.Product, err error) {
	table := "product"
	// Define the SQL query for searching products based on the provided criteria.
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

			((nullif($6, 0.00) is null or nullif($7, 0.00) is null) or ($6 <= shipping_price and $7 >= shipping_price)) and
			(nullif($6, 0.00) is null or $6 <= shipping_price) and
			(nullif($7, 0.00) is null or $7 >= shipping_price) and

			(($8::timestamp is null or $9::timestamp is null) or ($8::timestamp <= joined_at and $9::timestamp >= joined_at)) and
			($8::timestamp is null or $8::timestamp <= joined_at) and
			($9::timestamp is null or $9::timestamp >= joined_at) and

			(($10::timestamp is null or $11::timestamp is null) or ($10::timestamp <= delivered_at and $11::timestamp >= delivered_at)) and
			($10::timestamp is null or $10::timestamp <= delivered_at) and
			($11::timestamp is null or $11::timestamp >= delivered_at) and
			
			((nullif($12, 0) is null or nullif($13, 0) is null) or ($12 <= quantity and $13 >= quantity)) and
			(nullif($12, 0) is null or $12 <= quantity) and
			(nullif($13, 0) is null or $13 >= quantity) and
			client_id = $14
	`, table)

	// Execute the query with the provided search criteria and retrieve rows from the database.
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
		srch.QuantityRange.Start, srch.QuantityRange.End, srch.ClientID,
	)
	if err != nil {
		// If an error occurs while querying, wrap it with a meaningful error message and code.
		err = errorInRow(table, "get", err)
		return
	}

	// Initialize a slice to store the retrieved products.
	ps = make([]*product.Product, 0)
	// Iterate through each row of the result set.
	for rows.Next() {
		p := new(product.Product)
		// Scan the row's columns into the 'p' variable.
		err = rows.Scan(&p.ID, &p.ClientID, &p.GuideNumber, &p.Type, &p.JoinedAt, &p.DeliveredAt, &p.ShippingPrice, &p.VehiclePlate, &p.Port, &p.Vault, &p.Quantity)
		if err != nil {
			// If an error occurs during scanning, wrap it with additional error information.
			err = errorInRow(table, "scan", err)
			ps = nil
			return
		}

		// Append the scanned product to the 'ps' slice.
		ps = append(ps, p)
	}
	// Check for any error that occurred during iteration.
	err = rows.Err()
	if err != nil {
		// If an error occurred while iterating through rows, wrap it with additional error information.
		ps = nil
		err = errorInRows(table, "scanning", err)
	}
	return
}

// Update updates a product in the database.
func (pr ProductRepository) Update(p product.Product) (err error) {
	// Check if the user has ownership of the product before updating.
	err = pr.checkProductOwner(p.ID, p.ClientID)
	if err != nil {
		return
	}

	table := "product"
	// Define the SQL query for updating a product in the database.
	query := fmt.Sprintf(`
		update
			%s
		set
			guide_number = coalesce($1, guide_number),
			type = coalesce($2, type),
			joined_at = coalesce($3, joined_at),
			delivered_at = coalesce($4, delivered_at),
			shipping_price = coalesce($5, shipping_price),
			vehicle_plate = coalesce($6, vehicle_plate),
			port = coalesce($7, port),
			vault = coalesce($8, vault),
			quantity = coalesce($9, quantity)
		where
			id = $10
	`, table)

	// Execute the update query with the provided product details and ID.
	_, err = pr.db.Exec(query, &p.GuideNumber, &p.Type, &p.JoinedAt, &p.DeliveredAt, &p.ShippingPrice, &p.VehiclePlate, &p.Port, &p.Vault, &p.Quantity, p.ID)
	if err != nil {
		// If an error occurs during the update query, wrap it with additional error information.
		err = errorInRow(table, "update", err)
	}
	return
}

// Delete removes a product from the database.
func (pr ProductRepository) Delete(id, clientID int) (err error) {
	// Check if the user has ownership of the product before deleting.
	err = pr.checkProductOwner(id, clientID)
	if err != nil {
		return
	}

	table := "product"
	// Define the SQL query for deleting a product from the database.
	query := fmt.Sprintf(`
		delete from
			%s
		where
			id = $1
	`, table)

	// Execute the delete query with the provided product ID.
	_, err = pr.db.Exec(query, id)
	if err != nil {
		// If an error occurs during the delete query, wrap it with additional error information.
		err = errorInRow(table, "delete", err)
	}
	return
}

// checkProductOwner verifies if the user has ownership of the product with the given ID.
func (pr ProductRepository) checkProductOwner(id, clientID int) (err error) {
	table := "product"
	// Define the SQL query for checking product ownership by comparing the client ID.
	query := fmt.Sprintf(`
		select
			client_id = $2
		from
			%s
		where
			id = $1
	`, table)

	var isSame bool
	// Execute the query to check if the client ID matches the product's client ID.
	err = pr.db.QueryRow(query, id, clientID).Scan(&isSame)
	if err != nil {
		// If an error occurs during the query, wrap it with additional error information.
		err = errorInRow(table, "get", err)
		return
	}
	if !isSame {
		// If the client ID does not match, return an error indicating invalid ownership.
		err = fmt.Errorf("invalid client id: client id is not the same as the data to deal with")
	}
	return
}
