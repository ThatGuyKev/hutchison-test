package infrastructure

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"strings"

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

	// SEED DATA

	seedData := make(map[string][]string)

	file, err := os.Open("dogs.json")

	if err != nil {
		log.Printf("couldn't open seed file %v", err)
	} else {

		defer file.Close()

		if err := json.NewDecoder(file).Decode(&seedData); err != nil {
			log.Print("could read seed file")
		} else {
			for breed, variants := range seedData {
				if len(variants) < 1 {
					res, err := db.Exec("INSERT OR IGNORE INTO dogs (breed) VALUES (?)", breed)
					if err != nil {
						log.Printf("error inserting into dogs: %v", err)
					} else {
						log.Print(res)
					}
				} else {
					res, err := db.Exec("INSERT OR IGNORE INTO dogs (breed, variants) VALUES (?, json_array(?))", breed, strings.Join(variants, ", "))
					if err != nil {
						log.Printf("error inserting into dogs: %v", err)
					} else {
						log.Print(res)
					}
				}

			}
		}

	}

	return db
}
