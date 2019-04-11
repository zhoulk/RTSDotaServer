package entry

import (
	"server/define"
)

type Player struct {
	UserId   string
	Account  string
	Password string
	Name     string

	BaseInfo   *BaseInfo
	ExtendInfo *ExtendInfo
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
}

func NewBaseInfo() *BaseInfo {
	baseInfo := new(BaseInfo)
	baseInfo.Gold = 10000
	baseInfo.Diamond = 1000
	baseInfo.Level = 1
	baseInfo.Power = define.PLAYER_MAX_MP
	baseInfo.MaxPower = define.PLAYER_MAX_MP
	baseInfo.Exp = 0
	baseInfo.LevelUpExp = 90
	return baseInfo
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
}

func (p *Player) levelUp() {
	p.BaseInfo.Exp -= p.BaseInfo.LevelUpExp
	p.BaseInfo.Level += 1
	p.BaseInfo.LevelUpExp = playerExpList[p.BaseInfo.Level]
}
