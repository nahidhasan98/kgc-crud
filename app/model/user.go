package model

import (
	"errors"

	"github.com/nahidhasan98/kgc-crud/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

func getUserByUsername(username string) (*User, error) {
	// connecting to DB
	DB, ctx, cancel := config.DBConnect()
	defer cancel()
	defer DB.Client().Disconnect(ctx)

	// taking DB collection/table to a variable
	userCollection := DB.Collection("user")

	var dbUser User
	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&dbUser)
	if err != nil {
		return nil, err
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
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("username not found")
		}
		return nil, err
	}
	if !hashPasswordMatched(dbUser.Password, reqUser.Password) {
		return nil, errors.New("incorrect password")
	}

	return dbUser, nil
}
