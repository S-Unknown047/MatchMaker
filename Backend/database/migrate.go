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

	userInfoTable := `CREATE TABLE IF NOT EXISTS user_info (
	email TEXT PRIMARY KEY,
	user_name CITEXT UNIQUE NOT NULL,
	age INT,
	gender TEXT,
	level INT,
	profile_picture_url TEXT,
	profile_picture_id TEXT,
	playing_style TEXT[] DEFAULT '{}',                                                                  TEXT[] DEFAULT '{}',
	games TEXT[] DEFAULT '{}',
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`

	// 	gameTable := `CREATE TABLE IF NOT EXISTS joined_games (
	// 	game_id INT PRIMARY KEY,
	// 	game_name TEXT,
	// 	game_type TEXT,
	// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	// 	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	// );`

	fmt.Println("Database table created")
	if _, err := Conn.Exec(Ctx, table); err != nil {
		fmt.Println(err)
		return
	}

	if _, err := Conn.Exec(Ctx, userInfoTable); err != nil {
		fmt.Println(err)
	}

	// if _, err := Conn.Exec(Ctx, gameTable); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	fmt.Println("Database table created")

}
