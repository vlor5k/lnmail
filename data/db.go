package data

import (
	"fmt"

	sqlite "gosqlite.org"
)

var DB *sqlite.DB

func SetupDB() {
	db, err := sqlite.OpenWAL("app.db")

	if err != nil {
		fmt.Println("Failed to open sqlite database...")
	}

	DB = db
	fmt.Println("Database opened")
}

// db embeds *sql.DB — every database/sql method works.
