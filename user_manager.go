package main

import (
	"golang.org/x/crypto/bcrypt"
	"shorten-url/models"
)

func PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func CreateUser(email string, password string) []error {
	encryptedPassword, _ := PasswordEncrypt(password)
	db := DBManager.DB
	if err := db.Create(
		&models.User{
			Email:             email,
			EncryptedPassword: encryptedPassword,
		},
	).GetErrors(); err != nil {
		return err
	}
	return nil
}

func GetUser(email string) *models.User {
	var user models.User
	db := DBManager.DB
	err := db.First(&user, "email = ?", email).Error
	if err != nil {
		return nil
	}
	return &user
}
