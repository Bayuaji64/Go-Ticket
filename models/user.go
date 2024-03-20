package models

import (
	"database/sql"
	"errors"
	"time"

	"example.com/rest-api/db"
	"example.com/rest-api/helper"
)

type User struct {
	ID        int64
	Name      string
	Email     string    `binding:"required"`
	Password  string    `binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) Save() error {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"

	hashedPassword, err := helper.HashPassword(u.Password)
	if err != nil {
		return err
	}

	err = db.DB.QueryRow(query, u.Name, u.Email, hashedPassword).Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = $1"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedID int64
	var retrievedPasswordHash string

	err := row.Scan(&retrievedID, &retrievedPasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user not found")
		}
		return err
	}

	passwordIsValid := helper.CheckPasswordHash(u.Password, retrievedPasswordHash)

	if !passwordIsValid {
		return errors.New("credential not valid")
	}

	u.ID = retrievedID

	return nil
}
