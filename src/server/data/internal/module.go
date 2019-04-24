package internal

import (
	"server/base"
	"server/data/entry"

	"github.com/jinzhu/gorm"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
	players        map[string]*entry.Player
	playerHeros    map[string][]*entry.Hero
	heros          []*entry.Hero
	skills         []*entry.Skill
	items          []*entry.Item
	chapters       []*entry.Chapter
	playerChapters map[string][]*entry.Chapter
	guanKas        []*entry.GuanKa
	playerGuanKas  map[string][]*entry.GuanKa
	groups         []*entry.Group
	battleCache    map[string]*entry.BattleInfo

	db *gorm.DB
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	m.players = make(map[string]*entry.Player)
	m.playerHeros = make(map[string][]*entry.Hero)
	m.playerChapters = make(map[string][]*entry.Chapter)
	m.playerGuanKas = make(map[string][]*entry.GuanKa)
	m.heros = make([]*entry.Hero, 0)
	//m.heros = InitHeros()
	m.skills = InitSkills()
	m.items = InitItems()
	m.chapters = InitChapters()
	m.guanKas = InitGuanKas()
	m.battleCache = make(map[string]*entry.BattleInfo)

	m.PersistentData()
}

func (m *Module) OnDestroy() {

}

func (m *Module) EffectByEarn(player *entry.Player, earn *entry.Earn) {
	log.Debug("[Module EffectByEarn ] %v", earn)
	// 角色升级
	player.EffectByEarn(earn)
	// 英雄升级
	selelctHeros := m.SelectHeros(player)
	for _, hero := range selelctHeros {
		hero.EffectByEarn(earn)
	}
}

func (m *Module) EffectByExpend(player *entry.Player, expend *entry.Expend) {
	player.BaseInfo.Power -= expend.Power
}
