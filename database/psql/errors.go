package psql

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/test/database/errors"
	"github.com/lib/pq"
)

// parseErrorType identifies the error type based on the error details.
func parseErrorType(err error) (r string) {
	// Check if the error is a pq.Error (PostgreSQL-specific error).
	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code.Name() {
		case "unique_violation":
			r = errors.ALREADY_EXISTS // Set error type to ALREADY_EXISTS for unique violation.
		default:
			r = errors.UNKNOWN // Set error type to UNKNOWN for other PostgreSQL errors.
		}
	}
	// Check if the error is sql.ErrNoRows (indicating no rows found).
	if err == sql.ErrNoRows {
		r = errors.NOT_FOUND // Set error type to NOT_FOUND for no rows found.
	}
	return
}

// errorInRow generates a formatted error message for a single row operation failure.
func errorInRow(table, action string, err error) error {
	return errors.NewError(
		parseErrorType(err), // Get the appropriate error type based on the error.
		fmt.Sprintf("failed to %s a row in %s table", action, table), // Construct error message.
		err.Error(), // Include the original error content.
	)
}

// errorInRows generates a formatted error message for multiple rows operation failure.
func errorInRows(table, action string, err error) error {
	return errors.NewError(
		parseErrorType(err), // Get the appropriate error type based on the error.
		fmt.Sprintf("failed to %s rows in %s table", action, table), // Construct error message.
		err.Error(), // Include the original error content.
	)
}
