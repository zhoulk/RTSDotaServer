package internal

import (
	"math/rand"
	"server/data"
	"server/data/entry"
	"server/define"
	"server/msg"
	"server/tool"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func handleOwnHero(args []interface{}) {
	log.Debug("game handleOwnHero")
	//m := args[0].(*msg.HeroRequest)
	a := args[1].(gate.Agent)

	// // 输出收到的消息的内容
	log.Debug("user %v", a.UserData())
	player := a.UserData().(*entry.Player)

	response := new(msg.HeroOwnResponse)
	response.Code = msg.ResponseCode_SUCCESS

	heros := make([]*msg.Hero, 0)
	ownHeros := data.Module.AllOwnHeros(player)
	for _, v := range ownHeros {
		hero := ConverHeroToMsgHero(v)
		heros = append(heros, hero)
	}
	response.Heros = heros

	a.WriteMsg(response)
}

func handleHeroSkills(args []interface{}) {
	log.Debug("game handleHeroSkills")
	m := args[0].(*msg.HeroSkillsRequest)
	a := args[1].(gate.Agent)

	heroId := m.HeroId
	player := a.UserData().(*entry.Player)

	hero := data.Module.FindAHero(player, heroId)
	if hero == nil {
		response := new(msg.HeroSkillsResponse)
		response.Code = msg.ResponseCode_FAIL
		err := new(msg.Error)
		err.Code = define.HeroRandomLevelErr
		err.Msg = define.ERRMAP[err.Code]
		response.Err = err
		a.WriteMsg(response)
		return
	}

	skills := data.Module.FindHeroSkills(hero)

	response := new(msg.HeroSkillsResponse)
	response.Code = msg.ResponseCode_SUCCESS
	response.Skills = make([]*msg.Skill, 0)
	for _, sk := range skills {
		response.Skills = append(response.Skills, ConverSkillToMsgSkill(sk))
	}
	a.WriteMsg(response)
}

func handleHeroItems(args []interface{}) {
	log.Debug("game handleHeroItems")

}

func handleRandomHero(args []interface{}) {
	log.Debug("game handleRandomHero")

	m := args[0].(*msg.HeroRandomRequest)
	a := args[1].(gate.Agent)

	if a.UserData() == nil {
		log.Error("handleRandomHero error UserData is nil")
		return
	}

	response := new(msg.HeroRandomResponse)

	level := m.GetLevel()
	player := a.UserData().(*entry.Player)
	heros := data.Module.AllHeros()
	switch level {
	case msg.HeroRandomLevel_GOOD:
		if player.BaseInfo.Gold < 1000 {
			response.Code = msg.ResponseCode_FAIL
			err := new(msg.Error)
			err.Code = define.HeroRandomGoldErr
			err.Msg = define.ERRMAP[err.Code]
			response.Err = err
		} else {
			rd := rand.Intn(len(heros))
			hero := new(entry.Hero)
			tool.DeepCopy(hero, heros[rd])
			hero.HeroId = tool.UniqueId()
			hero.PlayerId = player.UserId
			data.Module.SavePlayerHero(player, hero)
			player.BaseInfo.Gold -= 1000

			response.Hero = ConverHeroToMsgHero(hero)
			response.Code = msg.ResponseCode_SUCCESS
		}
		break
	case msg.HeroRandomLevel_BETTER:
		if player.BaseInfo.Diamond < 20 {
			response.Code = msg.ResponseCode_FAIL
			err := new(msg.Error)
			err.Code = define.HeroRandomDiamondErr
			err.Msg = define.ERRMAP[err.Code]
			response.Err = err
		} else {
			rd := rand.Intn(len(heros))
			hero := new(entry.Hero)
			tool.DeepCopy(hero, heros[rd])
			hero.HeroId = tool.UniqueId()
			hero.PlayerId = player.UserId
			data.Module.SavePlayerHero(player, hero)
			player.BaseInfo.Diamond -= 20

			response.Hero = ConverHeroToMsgHero(hero)
			response.Code = msg.ResponseCode_SUCCESS
		}
		break
	case msg.HeroRandomLevel_BEST:
		if player.BaseInfo.Diamond < 200 {
			response.Code = msg.ResponseCode_FAIL
			err := new(msg.Error)
			err.Code = define.HeroRandomDiamondErr
			err.Msg = define.ERRMAP[err.Code]
			response.Err = err
		} else {
			rd := rand.Intn(len(heros))
			hero := new(entry.Hero)
			tool.DeepCopy(hero, heros[rd])
			hero.HeroId = tool.UniqueId()
			hero.PlayerId = player.UserId
			data.Module.SavePlayerHero(player, hero)
			player.BaseInfo.Diamond -= 200

			response.Hero = ConverHeroToMsgHero(hero)
			response.Code = msg.ResponseCode_SUCCESS
		}
		break
	default:
		response.Code = msg.ResponseCode_FAIL
		err := new(msg.Error)
		err.Code = define.HeroRandomLevelErr
		err.Msg = define.ERRMAP[err.Code]
		response.Err = err
	}

	a.WriteMsg(response)
}

func handleAllHero(args []interface{}) {
	log.Debug("game handleAllHero")

	//m := args[0].(*msg.HeroRequest)
	a := args[1].(gate.Agent)

	// // 输出收到的消息的内容
	log.Debug("user %v", a.UserData())

	response := new(msg.HeroResponse)
	response.Code = msg.ResponseCode_SUCCESS

	heros := make([]*msg.Hero, 0)
	for _, v := range data.Module.AllHeros() {
		hero := ConverHeroToMsgHero(v)
		heros = append(heros, hero)
	}
	response.Heros = heros

	a.WriteMsg(response)
}

func handleHeroSelect(args []interface{}) {
	log.Debug("game handleHeroSelect")
	m := args[0].(*msg.HeroSelectRequest)
	a := args[1].(gate.Agent)

	heroId := m.HeroId
	pos := m.Pos
	player := a.UserData().(*entry.Player)

	// 判断玩家是否拥有该英雄
	hero := data.Module.FindAHero(player, heroId)
	if hero == nil {
		response := new(msg.HeroSelectResponse)
		response.Code = msg.ResponseCode_FAIL
		response.Err = NewErr(define.HeroSelectExistErr)
		a.WriteMsg(response)
		return
	}
	// 判断位置是否有效
	if pos < 1 || pos > 9 {
		response := new(msg.HeroSelectResponse)
		response.Code = msg.ResponseCode_FAIL
		response.Err = NewErr(define.HeroSelectPosErr)
		a.WriteMsg(response)
		return
	}
	// 上阵
	// 1. 当前位置是否有英雄，如果有则先自动下阵
	oldHero := data.Module.FindAHeroAt(player, pos)
	if oldHero != nil {
		data.Module.UnSelectHero(player, oldHero)
	}
	// 2. 将目标英雄上阵
	data.Module.SelectHero(player, hero, pos)

	response := new(msg.HeroSelectResponse)
	response.Code = msg.ResponseCode_SUCCESS

	heroIds := data.Module.SelectHeroIds(player)
	response.HeroIds = heroIds

	a.WriteMsg(response)
}

func handleHeroUnSelect(args []interface{}) {
	log.Debug("game handleHeroUnSelect")
	m := args[0].(*msg.HeroUnSelectRequest)
	a := args[1].(gate.Agent)

	heroId := m.HeroId
	player := a.UserData().(*entry.Player)

	// 判断玩家是否拥有该英雄
	hero := data.Module.FindAHero(player, heroId)
	if hero == nil {
		response := new(msg.HeroUnSelectResponse)
		response.Code = msg.ResponseCode_FAIL
		response.Err = NewErr(define.HeroSelectExistErr)
		a.WriteMsg(response)
		return
	}
	// 下阵
	data.Module.UnSelectHero(player, hero)

	response := new(msg.HeroUnSelectResponse)
	response.Code = msg.ResponseCode_SUCCESS

	heroIds := data.Module.SelectHeroIds(player)
	response.HeroIds = heroIds

	a.WriteMsg(response)
}

func ConverHeroToMsgHero(v *entry.Hero) *msg.Hero {
	hero := new(msg.Hero)
	hero.Id = v.Id
	hero.Name = v.Name
	hero.Level = v.Level
	hero.Type = v.Type
	hero.Strength = v.Strength
	hero.StrengthStep = v.StrengthStep
	hero.Agility = v.Agility
	hero.AgilityStep = v.AgilityStep
	hero.Intelligence = v.Intelligence
	hero.IntelligenceStep = v.IntelligenceStep
	hero.Armor = v.Armor
	hero.AttackMin = v.AttackMin
	hero.AttackMax = v.AttackMax
	hero.Blood = v.Blood
	hero.SkillIds = make([]string, 0)
	if v.Skills != nil {
		for _, skill := range v.Skills {
			hero.SkillIds = append(hero.SkillIds, skill.SkillId)
		}
	}
	hero.ItemIds = make([]string, 0)
	hero.HeroId = v.HeroId
	hero.PlayerId = v.PlayerId
	hero.IsSelect = v.IsSelect
	hero.Pos = v.Pos
	hero.LevelUpExp = v.LevelUpExp
	hero.MP = v.MP
	hero.MaxMP = v.MaxMP
	hero.MaxBlood = v.MaxBlood
	hero.Exp = v.Exp
	return hero
}
