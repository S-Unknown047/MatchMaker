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
	var g game
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	res := middleware.SearchGame(g.GameName)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
