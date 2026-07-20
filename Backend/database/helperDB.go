package database

import (
	"errors"
	"matchMaker/model"
	"time"

	"github.com/jackc/pgx/v5"
)

func AddUser(user model.User) error {
	query := `INSERT INTO users (email, password, created_at, updated_at, roles) VALUES ($1, $2, $3, $4, $5)`
	if _, err := Conn.Exec(Ctx, query, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.Roles); err != nil {
		return err
	}

	return nil
}

func GetUser(Email string) (model.User, error) {
	query := `SELECT * FROM users WHERE email = $1`
	var data model.User
	err := Conn.QueryRow(Ctx, query, Email).Scan(
		&data.Email,
		&data.Password,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.RefreshToken,
		&data.Roles)

	if err != nil {
		return data, err
	}

	return data, nil
}

func UpdatePassword(user model.User) error {
	query := `UPDATE users SET password = $1, update_at = $2 WHERE email = $3`

	if _, err := Conn.Exec(Ctx, query, user.Password, user.UpdatedAt, user.Email); err != nil {
		return err
	}
	return nil
}

func SetRefreshToken(refreshToken string, user *model.User) error {
	query := `UPDATE users SET refresh_token = $1, updated_at = $2 WHERE email = $3`
	if _, err := Conn.Exec(Ctx, query, refreshToken, time.Now(), user.Email); err != nil {
		return err
	}
	return nil
}

func GetRefreshToken(refreshToken string) (model.User, error) {
	query := `SELECT email, password, created_at, updated_at, refresh_token FROM users WHERE refresh_token = $1`
	var data model.User
	if err := Conn.QueryRow(Ctx, query, refreshToken).Scan(
		&data.Email,
		&data.Password,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.RefreshToken); err != nil {
		return data, err
	}
	return data, nil
}

func SetUserData(data *model.UserProfile) error {
	query := `INSERT INTO user_info (user_name, age, gender, level, profile_picture_url, profile_picture_id, playing_style, games, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	if _, err := Conn.Exec(Ctx, query, data.UserName, data.Age, data.Gender, data.Level, data.ProfilePictureUrl, data.ProfilePictureId, data.PlayingStyle, data.Games, data.CreatedAt, data.UpdatedAt); err != nil {
		return err
	}
	return nil
}

func GetUserData(userName string) (model.UserProfile, error) {
	query := `SELECT * FROM user_info WHERE user_name = $1`
	var user model.UserProfile
	if err := Conn.QueryRow(Ctx, query, userName).Scan(
		&user.UserName,
		&user.Age,
		&user.Gender,
		&user.Level,
		&user.ProfilePictureUrl,
		&user.ProfilePictureId,
		&user.PlayingStyle,
		&user.Games,
		&user.CreatedAt,
		&user.UpdatedAt); err != nil {
		return user, err
	}

	return user, nil
}

func CheckUsername(userName string) error {
	query := `SELECT user_name FROM user_info WHERE user_name = $1`
	var user string
	err := Conn.QueryRow(Ctx, query, userName).Scan(&user)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		} else {
			return err
		}
	}

	return errors.New("Username already exists")
}

func UpdateUserProfile(data *model.UserProfile) error {
	query := `UPDATE user_info SET age = $1, gender = $2, level = $3, profile_picture_url = $4, profile_picture_id = $5, playing_style = $6, games = $7, updated_at = $8 WHERE user_name = $9`
	if _, err := Conn.Exec(Ctx, query, data.Age, data.Gender, data.Level, data.ProfilePictureUrl, data.ProfilePictureId, data.PlayingStyle, data.Games, data.UpdatedAt, data.UserName); err != nil {
		return err
	}
	return nil
}
