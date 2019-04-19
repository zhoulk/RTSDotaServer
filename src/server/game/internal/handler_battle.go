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

func handleBattleCreate(args []interface{}) {
	log.Debug("game handleBattleResult")

	m := args[0].(*msg.BattleCreateRequest)
	a := args[1].(gate.Agent)

	createType := m.GetType()
	createArgs := m.GetArgs()
	player := a.UserData().(*entry.Player)

	switch createType {
	case entry.BattleTypeGuanKa:
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
			// 是否开启
			guanKa := data.Module.FindGuanKa(player, int32(gkIdNum))
			if !guanKa.IsOpen {
				response := new(msg.BattleCreateResponse)
				response.Code = msg.ResponseCode_FAIL
				response.Err = NewErr(define.BattleGuanKaOpenErr)
				a.WriteMsg(response)
				return
			}
			// 次数是否充足
			if guanKa.Times <= 0 {
				response := new(msg.BattleCreateResponse)
				response.Code = msg.ResponseCode_FAIL
				response.Err = NewErr(define.BattleGuanKaTimesErr)
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
			// 是否上阵
			selectHeros := data.Module.SelectHeros(player)
			if selectHeros == nil || len(selectHeros) == 0 {
				response := new(msg.BattleCreateResponse)
				response.Code = msg.ResponseCode_FAIL
				response.Err = NewErr(define.BattleNoneHeroErr)
				a.WriteMsg(response)
				return
			}

			battleId := tool.UniqueId()
			data.Module.AddGuanKaBattle(battleId, guanKa)

			response := new(msg.BattleCreateResponse)
			response.Code = msg.ResponseCode_SUCCESS
			response.BattleId = battleId
			a.WriteMsg(response)
			return
		}
		break
	}
}

func handleBattleResult(args []interface{}) {
	log.Debug("game handleBattleResult")

	m := args[0].(*msg.BattleResultRequest)
	a := args[1].(gate.Agent)

	battleId := m.GetBattleId()
	result := m.GetResult()
	player := a.UserData().(*entry.Player)

	log.Debug("result ", result)

	battleInfo := data.Module.FindBattle(battleId)
	if battleInfo == nil {
		response := new(msg.BattleResultResponse)
		response.Code = msg.ResponseCode_FAIL
		response.Err = NewErr(define.BattleResultExistErr)
		a.WriteMsg(response)
		return
	}

	if battleInfo.Type == entry.BattleTypeGuanKa {
		earn := handleGuanKaBattle(result, battleInfo.Guanka)

		data.Module.EffectByEarn(player, earn)
		data.Module.EffectByExpend(player, battleInfo.Guanka.Expend)
		data.Module.UpdateGuanKa(player, battleInfo.Guanka, result)

		gkNotify := new(msg.GuanKaUpdateNotify)
		gkNotify.GuanKas = make([]*msg.GuanKa, 0)
		gkNotify.GuanKas = append(gkNotify.GuanKas, ConverGuanKaToMsgGuanKa(battleInfo.Guanka))
		nextGk := data.Module.FindNextGuanKa(player, battleInfo.Guanka.Id)
		if nextGk != nil {
			gkNotify.GuanKas = append(gkNotify.GuanKas, ConverGuanKaToMsgGuanKa(nextGk))
		}
		a.WriteMsg(gkNotify)

		response := new(msg.BattleResultResponse)
		response.Code = msg.ResponseCode_SUCCESS
		response.Earn = ConverEarnToMsgEarn(earn)
		response.Level = player.BaseInfo.Level
		response.Exp = player.BaseInfo.Exp
		response.LevelUpExp = player.BaseInfo.LevelUpExp
		a.WriteMsg(response)
	}
}

func handleGuanKaBattle(result int32, gk *entry.GuanKa) *entry.Earn {
	earn := new(entry.Earn)
	switch result {
	case entry.BattleResultStar1,
		entry.BattleResultStar2,
		entry.BattleResultStar3:
		earn.Gold = gk.Earn.Gold
		earn.HeroExp = gk.Earn.HeroExp
		earn.PlayerExp = gk.Earn.PlayerExp
		earn.ItemIds = make([]int32, 0)

		itemCnt := len(gk.Earn.ItemIds)

		indexArr := tool.C_M_N(int32(itemCnt), 1)

		for _, index := range indexArr {
			earn.ItemIds = append(earn.ItemIds, gk.Earn.ItemIds[index])
		}

		break
	case entry.BattleResultStar0:
		earn.Gold = 0
		earn.HeroExp = 0
		earn.PlayerExp = gk.Earn.PlayerExp
		earn.ItemIds = make([]int32, 0)
		break
	}
	return earn
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
