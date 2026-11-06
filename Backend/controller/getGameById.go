package controller

import (
	"encoding/json"
	"matchMaker/middleware"
	"net/http"
)

func GetGameByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		id = r.PathValue("id")
	}
	if id == "" && r.Body != nil {
		var bodyId string
		if err := json.NewDecoder(r.Body).Decode(&bodyId); err == nil {
			id = bodyId
		}
	}

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "id parameter is required",
		})
		return
	}

	response := middleware.GetGameByID(id)

	if response == nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "No Games found",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}