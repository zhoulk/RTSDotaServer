package internal

import (
	"server/data"
	"server/data/entry"
	"server/define"
	"server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func handleBattleGuanKa(args []interface{}) {
	log.Debug("game handleBattleGuanKa")

	m := args[0].(*msg.BattleGuanKaRequest)
	a := args[1].(gate.Agent)

	guanKaId := m.GetGuanKaId()
	player := a.UserData().(*entry.Player)

	// 判断玩家的该关卡是否开启
	guanKa := data.Module.FindGuanKa(player, guanKaId)
	if !guanKa.IsOpen {
		response := new(msg.BattleGuanKaResponse)
		response.Code = msg.ResponseCode_FAIL
		response.Err = NewErr(define.BattleGuanKaOpenErr)
		a.WriteMsg(response)
		return
	}
	// 判断体力是否充足
	if player.BaseInfo.Power < 6 {
		response := new(msg.BattleGuanKaResponse)
		response.Code = msg.ResponseCode_FAIL
		response.Err = NewErr(define.BattlePlayerPowerErr)
		a.WriteMsg(response)
		return
	}
	// 判断阵型是否有英雄
	heroIds := data.Module.SelectHeroIds(player)
	if len(heroIds) == 0 {
		response := new(msg.BattleGuanKaResponse)
		response.Code = msg.ResponseCode_FAIL
		response.Err = NewErr(define.BattleNoneHeroErr)
		a.WriteMsg(response)
		return
	}
	// 战斗
	res := fightGuanKa(player, guanKa)
	// 计算收益
	earn := CalGuanKaEarn(guanKa, res)
	// 收益生效
	data.Module.EffectByEarn(player, earn)
	data.Module.EffectByExpend(player, guanKa.Expend)

	response := new(msg.BattleGuanKaResponse)
	response.Code = msg.ResponseCode_SUCCESS

	response.Result = define.BattleResult_Success
	resGuanKa := ConverGuanKaToMsgGuanKa(guanKa)
	resGuanKa.Earn = ConverEarnToMsgEarn(earn)
	response.Guanka = resGuanKa

	a.WriteMsg(response)
}
