package controller

import (
	"encoding/json"
	"matchMaker/middleware"
	"net/http"
)

type game struct {
	GameName string `json:"gameName"`
}

func SearchGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	gameName := r.URL.Query().Get("game")

	if gameName == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "gameName parameter is required"})
		return
	}

	res := middleware.SearchGame(gameName)
	if res == nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed to search games"})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
