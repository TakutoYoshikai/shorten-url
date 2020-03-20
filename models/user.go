package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email             string `form:"email" binding:"required" gorm:"unique;not null"`
	EncryptedPassword string `form:"password" binding:"required"`
}
