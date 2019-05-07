package gate

import (
	"server/game"
	"server/login"
	"server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.LoginRequest{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.RegisteRequest{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.ZoneRequest{}, login.ChanRPC)

	msg.Processor.SetRouter(&msg.HeroRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.HeroRandomRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.HeroOwnRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.HeroLotteryRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.HeroSelectRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.HeroUnSelectRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.HeroSkillsRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.HeroEquipsRequest{}, game.ChanRPC)

	msg.Processor.SetRouter(&msg.SkillRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.SkillUpgradeRequest{}, game.ChanRPC)

	msg.Processor.SetRouter(&msg.ItemRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.EquipRequest{}, game.ChanRPC)

	msg.Processor.SetRouter(&msg.ChapterRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.GuanKaRequest{}, game.ChanRPC)

	msg.Processor.SetRouter(&msg.BattleGuanKaRequest{}, game.ChanRPC)

	msg.Processor.SetRouter(&msg.BattleStartRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.BattleResultRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.BattleCreateRequest{}, game.ChanRPC)

	msg.Processor.SetRouter(&msg.GroupOwnRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.GroupListRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.GroupCreateRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.GroupMembersRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.GroupApplyRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.GroupApplyMembersRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.GroupContriRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.GroupOperRequest{}, game.ChanRPC)

}
