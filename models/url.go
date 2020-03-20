package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/jinzhu/gorm"
)

type URL struct {
	gorm.Model
	SrcId  string `form:"src" binding:"required" gorm:"unique;not null"`
	DstUrl string `form:"dst" binding:"required" gorm:"not null"`
	UserId uint   `form:"userId"`
}

func (url URL) Validate(db *gorm.DB) {
	err := validation.Validate(url.DstUrl,
		validation.Required,
		is.URL,
	)
	if err != nil {
		db.AddError(err)
	}
}
