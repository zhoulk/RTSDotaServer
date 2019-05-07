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
	handler(&msg.HeroSelectRequest{}, handleHeroSelect)
	handler(&msg.HeroUnSelectRequest{}, handleHeroUnSelect)
	handler(&msg.HeroEquipsRequest{}, handleHeroEquips)
	handler(&msg.HeroSkillsRequest{}, handleHeroSkills)
	handler(&msg.HeroLotteryRequest{}, handleHeroLottery)

	handler(&msg.SkillRequest{}, handleAllSkill)
	handler(&msg.SkillUpgradeRequest{}, handleSkillUpgrade)

	handler(&msg.ItemRequest{}, handleAllItem)
	handler(&msg.EquipRequest{}, handleOwnEquip)

	handler(&msg.ChapterRequest{}, handleAllChapter)
	handler(&msg.GuanKaRequest{}, handleAllGuanKa)

	handler(&msg.BattleGuanKaRequest{}, handleBattleGuanKa)

	handler(&msg.BattleStartRequest{}, handleBattleStart)
	handler(&msg.BattleResultRequest{}, handleBattleResult)
	handler(&msg.BattleCreateRequest{}, handleBattleCreate)

	handler(&msg.GroupOwnRequest{}, handleGroupOwn)
	handler(&msg.GroupListRequest{}, handleGroupList)
	handler(&msg.GroupCreateRequest{}, handleGroupCreate)
	handler(&msg.GroupMembersRequest{}, handleGroupMembers)
	handler(&msg.GroupApplyRequest{}, handleGroupApply)
	handler(&msg.GroupApplyMembersRequest{}, handleGroupApplyMembers)
	handler(&msg.GroupOperRequest{}, handleGroupOper)
	handler(&msg.GroupContriRequest{}, handleGroupContri)

	handleCrontab()
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
