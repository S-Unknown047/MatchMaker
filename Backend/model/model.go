package model

import "time"

type User struct {
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	CreatedAt    time.Time `json:created_at`
	UpdatedAt    time.Time `json:updated_at`
	refreshToken string
}

type ReceivedSignupReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ReceivedLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
