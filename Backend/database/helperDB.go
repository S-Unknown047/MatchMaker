package database

import (
	"matchMaker/model"
	"time"
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
