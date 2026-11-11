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

	// rows, err := db.Query("")
	var lol int
	err = db.QueryRow("SELECT COUNT(*) FROM sqlite_master WHERE type='table';").Scan(&lol)
	if err != nil {
		panic("Couldn't query database")
	}

	// database already in place so nothing more to do
	if lol != 0 {
		return
	}

	MigrateDB()
}

func MigrateDB() {
	fmt.Println("Migrating...")
}
