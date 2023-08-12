package psql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import the PostgreSQL driver package (underscore indicates import for its side effects).
)

// properties holds the connection properties for the PostgreSQL database.
type properties struct {
	user string
	pass string
	name string
	host string
	port int
}

// PostgreSQLConnector is a struct representing a PostgreSQL database connector.
type PostgreSQLConnector struct {
	props properties // Connection properties
	db    *sql.DB    // Database connection instance
}

// Connect establishes a connection to the PostgreSQL database.
func (p *PostgreSQLConnector) Connect() (err error) {
	// Open a new database connection using the "postgres" driver and connection URL.
	db, err := sql.Open("postgres", connURL(p.props))
	if err != nil {
		return
	}

	// Ping the database to check if the connection is alive and working.
	err = db.Ping()
	if err != nil {
		err = fmt.Errorf("failed to ping database: %s", err)
		return
	}
	p.db = db
	return
}

// getConn returns the existing database connection or establishes a new one if not available.
func (p PostgreSQLConnector) getConn() (conn *sql.DB, err error) {
	if p.db == nil {
		err = p.Connect()
		if err != nil {
			return
		}
	}
	// Ping the database to ensure the connection is still alive.
	err = p.db.Ping()
	if err != nil {
		err = fmt.Errorf("failed to ping database: %s", err)
		return
	}

	conn = p.db
	return
}

// NewPostgreSQLConnector creates a new PostgreSQLConnector instance with the provided connection details.
func NewPostgreSQLConnector(user, pass, name, host string, port int) (conn *PostgreSQLConnector) {
	return &PostgreSQLConnector{
		props: properties{
			user: user,
			pass: pass,
			name: name,
			host: host,
			port: port,
		},
	}
}

// connURL generates the connection URL for the PostgreSQL database.
func connURL(props properties) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", props.user, props.pass, props.host, props.port, props.name, "disable")
}
