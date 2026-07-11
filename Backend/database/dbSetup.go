package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn
var Ctx = context.Background()

func Setup() {
	URL := os.Getenv("DATABASE_URL")
	fmt.Println(URL)
	Conn, err := pgx.Connect(Ctx, URL)

	if err != nil {
		Conn.Close(Ctx)
		log.Fatal(err)
	}

	CreateTable()
}
