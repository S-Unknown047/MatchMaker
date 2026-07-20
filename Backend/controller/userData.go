package controller

import (
	"encoding/json"
	"fmt"
	db "matchMaker/database"
	"matchMaker/middleware"
	"matchMaker/model"
	"net/http"
	"regexp"
	"time"
)

type reqUserData struct {
	UserName      string   `json:"user_name"`
	Age           int      `json:"age"`
	Gender        string   `json:"gender"`
	Level         int      `json:"level"`
	PlayingStyle  []string `json:"playing_style"`
	Games         []string `json:"games"`
	ProfilePiture []byte   `json:"profile_picture"`
}

var userNameRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{4,19}$`)

func SetUserData(w http.ResponseWriter, r *http.Request) {
	var data reqUserData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	ProfileId, secureURL := middleware.UploadImage(data.ProfilePiture)

	SendData := &model.UserProfile{
		UserName:          data.UserName,
		Age:               data.Age,
		Gender:            data.Gender,
		Level:             data.Level,
		PlayingStyle:      data.PlayingStyle,
		Games:             data.Games,
		ProfilePictureUrl: secureURL,
		ProfilePictureId:  ProfileId,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	err := db.SetUserData(SendData)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "ok",
		"data":  "User data set successfully",
	})
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("userName")

	if !userNameRegex.MatchString(userName) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid Username",
		})
		return
	}
	user, err := db.GetUserData(userName)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Internal Server Error",
		})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": "ok",
		"data":  user,
	})
}

func CheckUniqueUserName(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	if !userNameRegex.MatchString(username) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid Username",
		})
		return
	}

	if err := db.CheckUsername(username); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "ok",
		"data":  "Username is available",
	})
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var data reqUserData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	ProfileId, secureURL := middleware.UploadImage(data.ProfilePiture)

	SendData := &model.UserProfile{
		UserName:          data.UserName,
		Age:               data.Age,
		Gender:            data.Gender,
		Level:             data.Level,
		PlayingStyle:      data.PlayingStyle,
		Games:             data.Games,
		ProfilePictureUrl: secureURL,
		ProfilePictureId:  ProfileId,
		UpdatedAt:         time.Now(),
	}

	err := db.UpdateUserProfile(SendData)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "ok",
		"data":  "User data set successfully",
	})
}
