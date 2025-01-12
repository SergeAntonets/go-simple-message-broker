package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func ExecuteMigrations(db *sql.DB) {

	sql := `CREATE TABLE IF NOT EXISTS topics (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT);`

	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func Open() *sql.DB {
	db, err := sql.Open("sqlite3", "./message-broker.db")
	if err != nil {
		panic(err)
	}

	// defer db.Close()

	return db
}
