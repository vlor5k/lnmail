package data

import (
	"fmt"

	sqlite "gosqlite.org"
)

var DB *sqlite.DB

func SetupDB() {
	db, err := sqlite.OpenWAL("app.db")

	if err != nil {
		panic("Failed to open sqlite database...")
	}

	DB = db
	fmt.Println("Database opened")

	// rows, err := db.Query("SELECT COUNT(name) FROM sqlite_master WHERE type='table';")
	rows, err := db.Query("SELECT 1;")
	if err != nil {
		panic("Couldn't query database")
	}

	var lol int
	rows.Scan(&lol)

	fmt.Println(lol)
}

func MigrateDB() {

}
