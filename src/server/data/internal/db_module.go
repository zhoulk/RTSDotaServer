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

type UserBaseInfo struct {
	gorm.Model
	Uid        string `gorm:"size:64;unique;not null"`
	Name       string `gorm:"size:64"`
	Gold       int32
	Diamond    int32
	Exp        int32
	LevelUpExp int32
	Power      int32
	MaxPower   int32
	Level      int32
}
