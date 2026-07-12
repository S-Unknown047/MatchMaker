package controller

import (
	"encoding/json"
	db "matchMaker/database"
	"matchMaker/model"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var data model.ReceivedSignupReq
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}
	hashedpass, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}

	dbData := model.User{
		Email:     data.Email,
		Password:  string(hashedpass),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.AddUser(dbData); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "duplicate email entry",
				"status":  "error",
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "failed to create user",
			"status":  "error",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "user created successfully",
	})
}
