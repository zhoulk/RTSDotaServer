package entry

import (
	"server/define"
	"server/msg"
	"server/tool"
	"time"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

type Player struct {
	UserId     string
	Account    string
	Password   string
	Name       string
	LoginTime  time.Time
	LogoutTime time.Time

	BaseInfo   *BaseInfo
	ExtendInfo *ExtendInfo

	IsDirty bool

	Agent gate.Agent
}

func NewPlayer() *Player {
	player := new(Player)
	player.UserId = tool.UniqueId()
	player.IsDirty = true
	return player
}

func (p *Player) SetUserId(userId string) {
	p.UserId = userId
	p.IsDirty = true
}

func (p *Player) SetAccount(account string) {
	p.Account = account
	p.IsDirty = true
}

func (p *Player) SetPassword(password string) {
	p.Password = password
	p.IsDirty = true
}

func (p *Player) SetName(name string) {
	p.Name = name
	p.IsDirty = true
}

func (p *Player) SetLoginTime(loginTime time.Time) {
	p.LoginTime = loginTime
	p.IsDirty = true
}

func (p *Player) SetLogoutTime(logoutTime time.Time) {
	p.LogoutTime = logoutTime
	p.IsDirty = true
}

type BaseInfo struct {
	Gold       int32
	Diamond    int32
	Exp        int32
	LevelUpExp int32
	Power      int32
	MaxPower   int32
	Level      int32
	Military   int32

	IsDirty  bool
	IsNotify bool
}

func NewBaseInfo() *BaseInfo {
	baseInfo := new(BaseInfo)
	baseInfo.Gold = 10000
	baseInfo.Diamond = 5000
	baseInfo.Level = 1
	baseInfo.Power = define.PLAYER_MAX_MP
	baseInfo.MaxPower = define.PLAYER_MAX_MP
	baseInfo.Exp = 0
	baseInfo.LevelUpExp = 90

	baseInfo.IsDirty = true

	return baseInfo
}

func (b *BaseInfo) SetPower(power int32) {
	b.Power = power
	b.IsDirty = true
	b.IsNotify = true
}

func (b *BaseInfo) SetGold(gold int32) {
	log.Debug("SetGold  %v", gold)
	b.Gold = gold
	b.IsDirty = true
	b.IsNotify = true
}

func (b *BaseInfo) SetDiamond(diamond int32) {
	b.Diamond = diamond
	b.IsDirty = true
	b.IsNotify = true
}

func (b *BaseInfo) SetExp(exp int32) {
	b.Exp = exp
	b.IsDirty = true
	b.IsNotify = true
}

func (b *BaseInfo) SetLevel(level int32) {
	b.Level = level
	b.IsDirty = true
	b.IsNotify = true
}

func (b *BaseInfo) SetLevelUpExp(levelUpExp int32) {
	b.LevelUpExp = levelUpExp
	b.IsDirty = true
	b.IsNotify = true
}

type ExtendInfo struct {
	SelectHeroIds []string
	SelectHeros   []*Hero

	GroupId string

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

	IsDirty  bool
	IsNotify bool
}

func NewExtendInfo() *ExtendInfo {
	extendInfo := new(ExtendInfo)
	extendInfo.MaxFreeGoodLottery = 5
	extendInfo.MaxFreeBetterLottery = 1
	extendInfo.FreeGoodLottery = extendInfo.MaxFreeGoodLottery
	extendInfo.FreeBetterLottery = extendInfo.MaxFreeBetterLottery
	extendInfo.LastFreeGoodLotteryStamp = time.Now().Unix()
	extendInfo.LastFreeBetterLotteryStamp = time.Now().Unix()
	extendInfo.GoodLotteryCnt = 0
	extendInfo.BetterLotteryCnt = 0
	extendInfo.NeedGoodLotteryCnt = 10
	extendInfo.NeedBetterLotteryCnt = 10

	extendInfo.IsDirty = true

	return extendInfo
}

func (e *ExtendInfo) SetGroupId(groupId string) {
	e.GroupId = groupId
}

func (e *ExtendInfo) SetFreeGoodLottery(cnt int32) {
	e.FreeGoodLottery = cnt
	e.IsNotify = true
}

func (e *ExtendInfo) SetFreeBetterLottery(cnt int32) {
	e.FreeBetterLottery = cnt
	e.IsNotify = true
}

func (e *ExtendInfo) SetGoodLotteryCnt(cnt int32) {
	e.GoodLotteryCnt = cnt
	e.IsNotify = true
}

func (e *ExtendInfo) SetBetterLotteryCnt(cnt int32) {
	e.BetterLotteryCnt = cnt
	e.IsNotify = true
}

func ConverPlayerToMsgPlayer(v *Player) *msg.Player {
	player := new(msg.Player)
	player.UserId = v.UserId
	player.Name = v.Name
	if v.BaseInfo != nil {
		player.BaseInfo = ConverBaseInfoToMsgBaseInfo(v.BaseInfo)
	}
	if v.ExtendInfo != nil {
		player.ExtendInfo = ConverExtendInfoToMsgExtendInfo(v.ExtendInfo)
	}
	return player
}

func ConverBaseInfoToMsgBaseInfo(v *BaseInfo) *msg.BaseInfo {
	baseInfo := new(msg.BaseInfo)
	baseInfo.Level = v.Level
	baseInfo.Gold = v.Gold
	baseInfo.Diamond = v.Diamond
	baseInfo.Exp = v.Exp
	baseInfo.Power = v.Power
	baseInfo.LevelUpExp = v.LevelUpExp
	baseInfo.MaxPower = v.MaxPower
	return baseInfo
}

func ConverExtendInfoToMsgExtendInfo(v *ExtendInfo) *msg.ExtendInfo {
	extendInfo := new(msg.ExtendInfo)
	extendInfo.GroupId = v.GroupId
	return extendInfo
}
