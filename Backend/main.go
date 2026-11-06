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
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}
	fmt.Println("Server running on port:", Port)

	// Protected routes that require a valid JWT
	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("GET /games", controller.GetGames)
	protectedMux.HandleFunc("GET /searchgame", controller.SearchGame)
	protectedMux.HandleFunc("GET /getgamebyid", controller.GetGameByID)
	protectedMux.HandleFunc("GET /getgamebyid/{id}", controller.GetGameByID)
	protectedMux.HandleFunc("GET /checkuser/{name}", controller.CheckUniqueUserName)
	protectedMux.HandleFunc("GET /checkuser/", controller.CheckUniqueUserName)
	protectedMux.HandleFunc("GET /checkuser", controller.CheckUniqueUserName)
	protectedMux.HandleFunc("POST /upload", controller.SetUserData)
	protectedMux.HandleFunc("POST /setuserdata", controller.SetUserData)
	protectedMux.HandleFunc("GET /profile/{username}", controller.GetProfile)
	protectedMux.HandleFunc("POST /updateprofile", controller.UpdateProfile)

	// Root mux
	rootMux := http.NewServeMux()

	// Public routes (no JWT required)
	rootMux.HandleFunc("POST /login", controller.Login)
	rootMux.HandleFunc("POST /signup", controller.Signup)
	rootMux.HandleFunc("GET /refresh", controller.HndelRefreshToken)

	// Protected routes wrapped in JWT verification
	rootMux.Handle("/", middleware.JwtVerify(protectedMux))

	// CORS middleware as the outermost wrapper so OPTIONS requests and all endpoints have CORS
	handler := middleware.EnableCORS(rootMux)
	http.ListenAndServe(":"+Port, handler)
}
