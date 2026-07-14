package controller

import (
	"encoding/json"
	"matchMaker/middleware"
	"net/http"
)

func GetGames(w http.ResponseWriter, r *http.Request) {
	response := middleware.GetGames()
	if response == nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "No Games found"})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(string(response))
}
