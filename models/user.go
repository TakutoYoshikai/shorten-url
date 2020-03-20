package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email             string `form:"email" binding:"required" gorm:"unique;not null"`
	EncryptedPassword string `form:"password" binding:"required"`
}

func (user User) Validate(db *gorm.DB) {
	err := validation.Validate(user.Email,
		validation.Required,
		is.Email,
	)
	if err != nil {
		db.AddError(err)
	}
}
