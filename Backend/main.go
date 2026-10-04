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
	// server.HandleFunc("/ws")
	server.HandleFunc("POST /login", controller.Login)
	server.HandleFunc("POST /signup", controller.Signup)
	server.HandleFunc("GET /refresh", controller.HndelRefreshToken)
	server.HandleFunc("GET /games", controller.GetGames)
	server.HandleFunc("GET /searchgame", controller.SearchGame)
	server.HandleFunc("GET /checkuser/{name}", controller.CheckUniqueUserName)
	server.HandleFunc("GET /checkuser/", controller.CheckUniqueUserName)
	server.HandleFunc("POST /upload", controller.SetUserData)
	server.HandleFunc("POST /setuserdata", controller.SetUserData)
	server.HandleFunc("GET /profile/{username}", controller.GetProfile)
	server.HandleFunc("POST /updateprofile", controller.UpdateProfile)
	server.Handle("/", http.FileServer(http.Dir("../Frontend/dist")))

	handler := middleware.EnableCORS(server)
	http.ListenAndServe(":"+Port, handler)
}
