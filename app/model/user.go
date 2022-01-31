package model

import (
	"errors"

	"github.com/nahidhasan98/kgc-crud/config"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Username string
	Password string
}

func getUserByUsername(username string) (*User, error) {
	// taking DB
	DB := config.GetDB()
	defer DB.Close()

	rows, err := DB.Query(`SELECT id, username, password FROM user WHERE username = ?`, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dbUser User
	for rows.Next() {
		err = rows.Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password)
		if err != nil {
			return nil, err
		}
	}

	if dbUser.ID == 0 {
		return nil, errors.New("username not found")
	}

	return &dbUser, nil
}

func hashPasswordMatched(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Authenticate(reqUser *User) (*User, error) {
	dbUser, err := getUserByUsername(reqUser.Username)
	if err != nil {
		return nil, err
	}

	if !hashPasswordMatched(dbUser.Password, reqUser.Password) {
		return nil, errors.New("incorrect password")
	}

	return dbUser, nil
}
