package initdb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDatabase() (*sql.DB, error) {

	// Set up the database source string.
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "12345678", "127.0.0.1", 3306, "test2")

	// Create a database handle and open a connection pool.
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Check if our connection is alive.
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
