package internal

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Zone struct {
	ZoneId     string `gorm:"size:64;unique;not null"`
	TCPAddr    string `gorm:"size:32"`
	MaxConnNum int32
	Name       string `gorm:"size:64"`
	IsNew      bool

	gorm.Model
}

type User struct {
	Uid        string `gorm:"size:64;unique;not null"`
	Account    string `gorm:"size:128"`
	Password   string `gorm:"size:64"`
	LoginTime  time.Time
	LogoutTime time.Time

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
	Military   int32

	gorm.Model
}

type UserTarvern struct {
	Uid                        string `gorm:"size:64;unique;not null"`
	FreeGoodLottery            int32
	FreeBetterLottery          int32
	MaxFreeGoodLottery         int32
	MaxFreeBetterLottery       int32
	LastFreeGoodLotteryStamp   int64
	LastFreeBetterLotteryStamp int64
	GoodLotteryCnt             int32
	BetterLotteryCnt           int32
	NeedGoodLotteryCnt         int32
	NeedBetterLotteryCnt       int32

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

type GuanKaExpendDefine struct {
	GuanKaId int32 `gorm:"unique;not null"`
	Power    int32

	gorm.Model
}

type GuanKaEarnDefine struct {
	GuanKaId  int32  `gorm:"unique;not null"`
	ItemIds   string `gorm:"size:128"`
	HeroExp   int32
	PlayerExp int32
	Gold      int32

	gorm.Model
}

type GuanKaHeroDefine struct {
	GuanKaId int32  `gorm:"not null"`
	HeroId   string `gorm:"size:128"`

	gorm.Model
}

type EquipDefine struct {
	ItemId int32  `gorm:"unique;not null"`
	Name   string `gorm:"size:16"`
	Price  int32
	Effect string `gorm:"size:256"`
	Desc   string `gorm:"size:256"`

	gorm.Model
}

type EquipMixDefine struct {
	ItemId  int32
	ChildId int32
	Num     int32

	gorm.Model
}

type ConsumeDefine struct {
	ItemId int32  `gorm:"unique;not null"`
	Name   string `gorm:"size:16"`
	Price  int32
	Desc   string `gorm:"size:256"`

	gorm.Model
}

type HeroChipDefine struct {
	ItemId     int32  `gorm:"unique;not null"`
	Name       string `gorm:"size:16"`
	Price      int32
	ComposeCnt int32
	Desc       string `gorm:"size:256"`

	gorm.Model
}

type ExpHeroDefine struct {
	Level int32 `gorm:"unique;not null"`
	Exp   int32

	gorm.Model
}

type ExpPlayerDefine struct {
	Level int32 `gorm:"unique;not null"`
	Exp   int32

	gorm.Model
}

type UserHero struct {
	Uid              string `gorm:"size:64;not null"`
	HeroId           string `gorm:"size:64;unique;not null"`
	HeroDefineId     int32
	Level            int32
	Exp              int32
	LevelUpExp       int32
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
	IsSelect         bool
	Pos              int32

	gorm.Model
}

type UserHeroSkill struct {
	HeroId        string `gorm:"size:64;not null"`
	SkillId       string `gorm:"size:64;unique;not null"`
	SkillDefineId int32
	Level         int32
	IsOpen        bool

	gorm.Model
}

type UserChapter struct {
	Uid             string `gorm:"size:64;not null"`
	ChapterId       string `gorm:"size:64;unique;not null"`
	ChapterDefineId int32
	Star            int32
	Status          int32
	IsOpen          bool

	gorm.Model
}

type UserGuanKa struct {
	Uid            string `gorm:"size:64;not null"`
	GuanKaId       string `gorm:"size:64;unique;not null"`
	GuanKaDefineId int32
	Star           int32
	Status         int32
	IsOpen         bool
	Times          int32

	gorm.Model
}

type Group struct {
	GroupId          string `gorm:"size:64;unique;not null"`
	GroupName        string `gorm:"size:16"`
	GroupLeader      string `gorm:"size:64"`
	GroupDeclaration string `gorm:"size:128"`
	MemberCnt        int32
	MemberTotal      int32
	GroupLevel       int32
	ContriCurrent    int32
	ContriLevelUp    int32

	gorm.Model
}

type GroupMember struct {
	GroupId     string `gorm:"size:64;not null"`
	UserId      string `gorm:"size:64;unique;not null"`
	ContriToday int32
	ContriTotal int32
	Job         int32

	gorm.Model
}

type UserEquip struct {
	Uid          string `gorm:"size:64;not null"`
	EquipId      string `gorm:"size:64;unique;not null"`
	ItemDefineId int32  `gorm:"not null"`
	HeroId       string `gorm:"size:64"`

	gorm.Model
}

type UserConsume struct {
	Uid          string `gorm:"size:64;not null"`
	ConsumeId    string `gorm:"size:64;unique;not null"`
	ItemDefineId int32  `gorm:"not null"`

	gorm.Model
}

type UserHeroChip struct {
	Uid          string `gorm:"size:64;not null"`
	ChipId       string `gorm:"size:64;unique;not null"`
	ItemDefineId int32  `gorm:"not null"`
	Cnt          int32

	gorm.Model
}

type GameConfig struct {
	ConfigKey   string `gorm:"size:16;unique;not null"`
	ConfigValue string `gorm:"size:64"`

	gorm.Model
}
