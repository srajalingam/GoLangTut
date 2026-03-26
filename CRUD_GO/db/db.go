package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")
	if err != nil {
		panic("could not connect to DB")
	}

	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)

	createTables()
}

func createTables() {

	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL 
		)
	`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		fmt.Println(err)
		panic("Could not create users table")
	}

	createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        date_time TEXT NOT NULL,
        user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
    );
`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic(err)
	}
}
