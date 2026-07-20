package main

import (
	"fmt"
	"matchMaker/controller"
	db "matchMaker/database"
	"matchMaker/middleware"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	db.Setup()
	middleware.CreateInstance()
}

func main() {
	server := http.NewServeMux()
	Port := os.Getenv("PORT")
	fmt.Println(Port)
	server.HandleFunc("POST /login", controller.Login)
	server.HandleFunc("POST /signup", controller.Signup)
	server.HandleFunc("GET /refresh", controller.HndelRefreshToken)
	server.HandleFunc("GET /games", controller.GetGames)
	server.HandleFunc("GET /searchgame", controller.SearchGame)
	server.Handle("/", http.FileServer(http.Dir("../Frontend/dist")))

	http.ListenAndServe(":"+Port, server)
}
