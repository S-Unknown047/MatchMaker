package database

import (
	"fmt"
)

func CreateTable() {
	fmt.Println("Database table created")

	table := `CREATE TABLE IF NOT EXISTS users (
    email TEXT PRIMARY KEY,
    password TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    refresh_token TEXT DEFAULT '',
	roles TEXT[] DEFAULT '{user}'
);`
	fmt.Println("Database table created")

	if _, err := Conn.Exec(Ctx, table); err != nil {
		fmt.Println("here")
		fmt.Println(err)
		return
	}

	fmt.Println("Database table created")

}
