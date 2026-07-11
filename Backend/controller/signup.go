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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hashedpass, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	dbData := model.User{
		Email:     data.Email,
		Password:  string(hashedpass),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.AddUser(dbData); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("duplicate email entry"))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to create user"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user created successfully"))
}
