package internal

import (
	"server/data"
	"server/data/entry"
	"server/define"
	"server/msg"
	"server/tool"
	"strconv"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

const (
	BattleTypeGuanKa int32 = 1
)

func handleBattleCreate(args []interface{}) {
	log.Debug("game handleBattleResult")

	m := args[0].(*msg.BattleCreateRequest)
	a := args[1].(gate.Agent)

	// battleId := m.GetBattleId()
	createType := m.GetType()
	createArgs := m.GetArgs()
	player := a.UserData().(*entry.Player)

	switch createType {
	case BattleTypeGuanKa:
		if createArgs == nil || len(createArgs) < 1 {
			response := new(msg.BattleCreateResponse)
			response.Code = msg.ResponseCode_FAIL
			response.Err = NewErr(define.SysRequestArgsErr)
			a.WriteMsg(response)
			return
		} else {
			gkId := createArgs[0]
			gkIdNum, err := strconv.ParseInt(gkId, 0, 32)
			if err != nil {
				response := new(msg.BattleCreateResponse)
				response.Code = msg.ResponseCode_FAIL
				response.Err = NewErr(define.SysRequestArgsErr)
				a.WriteMsg(response)
				return
			}
			guanKa := data.Module.FindGuanKa(player, int32(gkIdNum))
			if !guanKa.IsOpen {
				response := new(msg.BattleCreateResponse)
				response.Code = msg.ResponseCode_FAIL
				response.Err = NewErr(define.BattleGuanKaOpenErr)
				a.WriteMsg(response)
				return
			}
			// 判断体力是否充足
			if player.BaseInfo.Power < 6 {
				response := new(msg.BattleCreateResponse)
				response.Code = msg.ResponseCode_FAIL
				response.Err = NewErr(define.BattlePlayerPowerErr)
				a.WriteMsg(response)
				return
			}
			selectHeros := data.Module.SelectHeros(player)
			if selectHeros == nil || len(selectHeros) == 0 {
				response := new(msg.BattleCreateResponse)
				response.Code = msg.ResponseCode_FAIL
				response.Err = NewErr(define.BattleNoneHeroErr)
				a.WriteMsg(response)
				return
			}
		}
		break
	}

	response := new(msg.BattleCreateResponse)
	response.Code = msg.ResponseCode_SUCCESS
	response.BattleId = tool.UniqueId()
	a.WriteMsg(response)
}

func handleBattleResult(args []interface{}) {
	log.Debug("game handleBattleResult")

	m := args[0].(*msg.BattleResultRequest)
	a := args[1].(gate.Agent)

	// battleId := m.GetBattleId()
	result := m.GetResult()
	player := a.UserData().(*entry.Player)

	log.Debug("result ", result)

	earn := new(entry.Earn)
	switch result {
	case entry.BattleResultStar1:
		earn.Gold = 100
		earn.HeroExp = 120
		earn.PlayerExp = 100
		earn.ItemIds = make([]int32, 0)
		earn.ItemIds = append(earn.ItemIds, 1)
		earn.ItemIds = append(earn.ItemIds, 2)
		break
	case entry.BattleResultStar2:
	case entry.BattleResultStar3:
		break
	case entry.BattleResultStar0:
		earn.Gold = 0
		earn.HeroExp = 0
		earn.PlayerExp = 100
		break
	}

	data.Module.EffectByEarn(player, earn)

	response := new(msg.BattleResultResponse)
	response.Code = msg.ResponseCode_SUCCESS
	response.Earn = ConverEarnToMsgEarn(earn)
	a.WriteMsg(response)
}

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
