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
type HeroDefine struct {
	gorm.Model
	HeroId           int32  `gorm:"unique;not null"`
	Name             string `gorm:"size:16"`
	Type             int32
	Strength         int32
	StrengthStep     int32
	Agility          int32
	AgilityStep      int32
	Intelligence     int32
	IntelligenceStep int32
	Armor            int32
	AttackMin        int32
	AttackMax        int32
	Blood            int32
	MP               int32
	SkillIds         string `gorm:"size:32"`
}
