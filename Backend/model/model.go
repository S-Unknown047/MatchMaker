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
	Email             string    `json:"email"`
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

type MsgType struct {
	channel_id string    `json:"channel_id"`
	Message    string    `json:"message"`
	Created    time.Time `json:"created"`
}

type Event struct {
	MsgType MsgType `json:"msg_type"`
}
type ChatReceived struct {
}

type GameCover struct {
	ID      int    `json:"id"`
	ImageID string `json:"image_id"`
	URL     string `json:"url"`
}

type Game struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Cover      GameCover `json:"cover"`
	Categories []string  `json:"categories"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
