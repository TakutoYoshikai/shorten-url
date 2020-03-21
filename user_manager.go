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

func CreateUser(email string, password string) (*models.User, []error) {
	encryptedPassword, _ := PasswordEncrypt(password)
	db := DBManager.DB
	user := models.User{
		Email:             email,
		EncryptedPassword: encryptedPassword,
	}
	if errs := db.Create(
		&user,
	).GetErrors(); errs != nil && len(errs) > 0 {
		return nil, errs
	}
	return &user, nil
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

func Login(email string, password string) *models.User {
	user := GetUser(email)
	if user == nil {
		return nil
	}
	dbPassword := user.EncryptedPassword
	if err := CompareHashAndPassword(dbPassword, password); err != nil {
		return nil
	}
	return user
}

func AllURL(user *models.User) []models.URL {
	var urls []models.URL
	db := DBManager.DB
	db.Where("user_id = ?", int(user.ID)).Find(&urls)
	return urls
}
