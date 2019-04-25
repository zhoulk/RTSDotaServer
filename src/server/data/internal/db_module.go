package internal

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Uid      string `gorm:"size:64;unique;not null"`
	Account  string `gorm:"size:128"`
	Password string `gorm:"size:64"`

	gorm.Model
}

type UserBaseInfo struct {
	Uid        string `gorm:"size:64;unique;not null"`
	Name       string `gorm:"size:64"`
	Gold       int32
	Diamond    int32
	Exp        int32
	LevelUpExp int32
	Power      int32
	MaxPower   int32
	Level      int32

	gorm.Model
}

type HeroDefine struct {
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

	gorm.Model
}

type SkillDefine struct {
	SkillId   int32  `gorm:"unique;not null"`
	Name      string `gorm:"size:16"`
	Type      int32
	LevelDesc string `gorm:"size:256"`
	Param1    string `gorm:"size:32"`
	Param2    string `gorm:"size:32"`
	Param3    string `gorm:"size:32"`
	Param4    string `gorm:"size:32"`
	Param5    string `gorm:"size:32"`
	Desc      string `gorm:"size:256"`

	gorm.Model
}

type ChapterDefine struct {
	ChapterId int32  `gorm:"unique;not null"`
	Name      string `gorm:"size:16"`
	GuanKaNum int32

	gorm.Model
}

type GuanKaDefine struct {
	GuanKaId   int32  `gorm:"unique;not null"`
	Name       string `gorm:"size:16"`
	ChapterId  int32
	TotalTimes int32

	gorm.Model
}

type ItemDefine struct {
	ItemId int32  `gorm:"unique;not null"`
	Name   string `gorm:"size:16"`
	Price  int32
	Effect string `gorm:"size:256"`
	Desc   string `gorm:"size:256"`

	gorm.Model
}

type ItemMixDefine struct {
	ItemId  int32
	ChildId int32
	Num     int32

	gorm.Model
}
