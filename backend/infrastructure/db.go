package infrastructure

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DBAdapter *sql.DB

func InitialiseDB() *sql.DB {

	db, err := sql.Open("sqlite", "app.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS dogs (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	breed TEXT NOT NULL UNIQUE,
	variants TEXT,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	deleted_at DATETIME 

	)`)
	if err != nil {
		panic(err)
	}

	DBAdapter = db
	return db
}
