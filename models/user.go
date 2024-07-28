package models

import (
	"errors"
	"events-mgmt-portal/db"
	"events-mgmt-portal/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

var users = []User{}

func (u User) Save() error {
	query := `INSERT INTO users(email, password) 
  VALUES(?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Invalid username or password")
	}

	isPasswordValid := utils.CheckHashedPassword(u.Password, retrievedPassword)

	if !isPasswordValid {
		return errors.New("Invalid username or password")
	}
	return nil
}
