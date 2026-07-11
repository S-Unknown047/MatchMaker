package main

import (
	"matchMaker/controller"
	db "matchMaker/database"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	db.Setup()
}
func main() {
	server := http.NewServeMux()
	Port := os.Getenv("PORT")
	server.HandleFunc("POST /login", controller.Login)
	server.HandleFunc("POST /signup", controller.Signup)
	server.HandleFunc("POST /refresh", controller.HndelRefreshToken)
	server.Handle("/", http.FileServer(http.Dir("../Frontend/dist")))

	http.ListenAndServe(":"+Port, server)
}
