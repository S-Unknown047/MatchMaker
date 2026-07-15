package controller

import (
	"encoding/json"
	"matchMaker/middleware"
	"net/http"
)

func GetGameByID(w http.ResponseWriter, r http.Request) {
	var id string
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error()})
		return
	}

	response := middleware.GetGameByID(id)

	if response == nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "No Games found"})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
