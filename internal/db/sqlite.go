package db

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
)

// https://www.sqlitetutorial.net/sqlite-go/create-table/
func InitDB() {
	// Connect to the SQLite database
	db, err := sql.Open("sqlite", "./data/my.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()
	fmt.Println("Connected to the SQLite database successfully.")

	// Get the version of SQLite
	var sqliteVersion string
	err = db.QueryRow("select sqlite_version()").Scan(&sqliteVersion)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sqliteVersion)
}
