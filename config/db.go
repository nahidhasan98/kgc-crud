package config

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// GetDB function for connenting to DB
func GetDB() *sql.DB {
	dbName := databaseName
	dbConnectionString := databaseConnectionString

	db, err := sql.Open(dbConnectionString, "../"+dbName)
	if err != nil {
		fmt.Println("database connection error")
		panic(err)
	}

	return db
}

// DBInit function for creating table to DB
func DBInit() {
	db := GetDB()
	defer db.Close()

	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, username TEXT, password TEXT)")
	if err != nil {
		fmt.Println("database initialization error")
		panic(err)
	}
	statement.Exec()

	statement, err = db.Prepare("CREATE TABLE IF NOT EXISTS student (id INTEGER PRIMARY KEY, name TEXT, email TEXT, phone TEXT, reg TEXT, session TEXT, roll TEXT, passingYear TEXT, avatarURL TEXT)")
	if err != nil {
		fmt.Println("database initialization error")
		panic(err)
	}
	statement.Exec()
}
