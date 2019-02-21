package Model

import (
	"github.com/jinzhu/gorm"
)

type Document struct {
	gorm.Model
	Code                          string `gorm:"type:varchar(20);not null"`
	FileName					  string `gorm:"type:varchar(150);not null"`
	UserId					  	  uint `gorm:"type:int();not null"`
}
