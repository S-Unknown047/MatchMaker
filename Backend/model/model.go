package model

import "time"

type User struct {
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	CreatedAt    time.Time `json:created_at`
	UpdatedAt    time.Time `json:updated_at`
	RefreshToken string    `json:"refresh_token"`
	Roles        []string  `json:"roles"`
}

type ReceivedSignupReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ReceivedLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserProfile struct {
	UserName          string    `json:"user_name"`
	Age               int       `json:"age"`
	Gender            string    `json:"gender"`
	Level             int       `json:"level"`
	ProfilePictureUrl string    `json:"profile_picture_url"`
	ProfilePictureId  string    `json:"profile_picture_id"`
	PlayingStyle      []string  `json:"playing_style"`
	Games             []string  `json:"games"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
