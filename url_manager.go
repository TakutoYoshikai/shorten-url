package main

import (
	"math/rand"
	"shorten-url/models"
)

func RandomId() string {
	return RandString(5)
}

var rs1Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}
	return string(b)
}

func CreateURL(user *models.User, dstUrl string) (*models.URL, []error) {
	db := DBManager.DB
	var userId int = -1
	if user != nil {
		userId = int(user.ID)
	}
	url := models.URL{
		UserId: userId,
		DstUrl: dstUrl,
		SrcId:  RandomId(),
	}
	if errs := db.Create(
		&url,
	).GetErrors(); errs != nil && len(errs) > 0 {
		return nil, errs
	}
	return &url, nil
}

func GetURL(srcId string) *models.URL {
	var url models.URL
	db := DBManager.DB
	err := db.First(&url, "src_id = ?", srcId).Error
	if err != nil {
		return nil
	}
	return &url
}
