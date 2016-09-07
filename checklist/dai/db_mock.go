// Package dai contains database mock for example.
// You can have a real database connection here.
package dai

import "github.com/jimmy-go/pgwp"

var (
	// Db database connection.
	Db *pgwp.Pool
)

// Configure prepares database.
func Configure(driver, connectURL string) error {
	var err error
	// Db, err = pgwp.Connect(driver, connectURL, 5, 5)
	return err
}

// Close closes database connections.
func Close() {
	Db.Close()
}
