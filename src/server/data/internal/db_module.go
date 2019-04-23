package internal

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Uid      string `gorm:"size:64;unique;not null"`
	Account  string `gorm:"size:128"`
	Password string `gorm:"size:64"`
}
