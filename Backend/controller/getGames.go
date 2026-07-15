package controller

import (
	"encoding/json"
	"fmt"
	"matchMaker/middleware"
	"net/http"
)

func GetGames(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	response := middleware.GetGames(page)
	w.Header().Set("Content-Type", "application/json")
	if response == nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "No Games found"})
		return
	}
	fmt.Println(string(response))
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
