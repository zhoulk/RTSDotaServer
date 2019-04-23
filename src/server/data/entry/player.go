package entry

import (
	"server/define"
	"server/tool"
	"time"
)

type Player struct {
	UserId   string
	Account  string
	Password string
	Name     string

	BaseInfo   *BaseInfo
	ExtendInfo *ExtendInfo

	IsDirty bool
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

func (p *Player) SetBaseInfo(baseInfo *BaseInfo) {
	p.BaseInfo = baseInfo
	p.IsDirty = true
}

func (p *Player) SetExtendInfo(extendInfo *ExtendInfo) {
	p.ExtendInfo = extendInfo
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
}

func NewPlayer() *Player {
	player := new(Player)
	player.UserId = tool.UniqueId()
	return player
}

func NewBaseInfo() *BaseInfo {
	baseInfo := new(BaseInfo)
	baseInfo.Gold = 10000
	baseInfo.Diamond = 0
	baseInfo.Level = 1
	baseInfo.Power = define.PLAYER_MAX_MP
	baseInfo.MaxPower = define.PLAYER_MAX_MP
	baseInfo.Exp = 0
	baseInfo.LevelUpExp = 90
	return baseInfo
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
	return extendInfo
}

func (p *Player) EffectByEarn(earn *Earn) {
	p.BaseInfo.Exp += earn.PlayerExp
	for {
		if p.BaseInfo.Exp < p.BaseInfo.LevelUpExp {
			break
		}
		if p.BaseInfo.Level+1 >= int32(len(playerExpList)) {
			break
		}
		p.levelUp()
	}

	p.IsDirty = true
}

func (p *Player) levelUp() {
	p.BaseInfo.Exp -= p.BaseInfo.LevelUpExp
	p.BaseInfo.Level += 1
	p.BaseInfo.LevelUpExp = playerExpList[p.BaseInfo.Level]
}
