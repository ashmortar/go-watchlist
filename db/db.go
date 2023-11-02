package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

// var dbUrl = "file:path/to/file.db"
// db, err := sql.Open("libsql", dbUrl)
// if err != nil {
//     fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbUrl, err)
//     os.Exit(1)
// }

func Open() (*sql.DB, error) {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return nil, fmt.Errorf("DB_URL is not set")
	}

	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to open db %s: %s", dbUrl, err)
	}

	return db, nil
}

func Close(db *sql.DB) error {
	return db.Close()
}
