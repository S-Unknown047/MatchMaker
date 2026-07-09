package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}
func main() {
	server := http.NewServeMux()
	Port := os.Getenv("PORT")
	server.Handle("/", http.FileServer(http.Dir("../Frontend/dist")))

	http.ListenAndServe(":"+Port, server)
}
