package internal

import (
	"reflect"
	"server/define"
	"server/msg"
)

func init() {
	handler(&msg.HeroRequest{}, handleAllHero)
	handler(&msg.HeroRandomRequest{}, handleRandomHero)
	handler(&msg.HeroOwnRequest{}, handleOwnHero)

	handler(&msg.SkillRequest{}, handleAllSkill)
	handler(&msg.ItemRequest{}, handleAllItem)
	handler(&msg.ChapterRequest{}, handleAllChapter)
	handler(&msg.GuanKaRequest{}, handleAllGuanKa)
	handler(&msg.BattleGuanKaRequest{}, handleBattleGuanKa)
	handler(&msg.HeroSelectRequest{}, handleHeroSelect)
	handler(&msg.HeroUnSelectRequest{}, handleHeroUnSelect)
	handler(&msg.SkillUpgradeRequest{}, handleSkillUpgrade)
	handler(&msg.HeroSkillsRequest{}, handleHeroSkills)
	handler(&msg.HeroItemsRequest{}, handleHeroItems)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func NewErr(errCode int32) *msg.Error {
	err := new(msg.Error)
	err.Code = errCode
	err.Msg = define.ERRMAP[err.Code]
	return err
}
