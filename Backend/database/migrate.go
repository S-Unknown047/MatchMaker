package database

import (
	"fmt"
)

func CreateTable() {
	table := `CREATE TABLE IF NOT EXISTS users (
    email TEXT PRIMARY KEY,
    password TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
    refresh_token TEXT DEFAULT '',
);`
	if _, err := Conn.Exec(Ctx, table); err != nil {
		fmt.Println(err)
		return
	}
}
