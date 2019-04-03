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
	log.Debug("[handler_battle ] fightGuanKa guanKaIs = %d %v", guanKa.Id, res)
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

func handleBattleStart(args []interface{}) {
	log.Debug("game handleBattleStart")

	// m := args[0].(*msg.BattleStartRequest)
	a := args[1].(gate.Agent)

	// battleId := m.GetBattleId()
	// player := a.UserData().(*entry.Player)

	response := new(msg.BattleStartResponse)
	response.Code = msg.ResponseCode_SUCCESS

	response.Heros = make([]*msg.Hero, 0)
	response.Skills = make([]*msg.Skill, 0)
	response.Items = make([]*msg.Item, 0)
	selectHeros := data.Module.AllHeros()
	// selectSkills := data.Module.AllSkills()
	// selectItems := data.Module.AllItems()
	for _, hero := range selectHeros {
		response.Heros = append(response.Heros, ConverHeroToMsgHero(hero))
	}
	// for _, skill := range selectSkills {
	// 	response.Skills = append(response.Skills, ConverSkillToMsgSkill(skill))
	// }
	// for _, item := range selectItems {
	// 	response.Items = append(response.Items, ConverItemToMsgItem(item))
	// }

	a.WriteMsg(response)
}
