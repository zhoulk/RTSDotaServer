package internal

import (
	"math/rand"
	"server/data"
	"server/data/entry"
	"server/define"
	"server/msg"
	"server/tool"
	"time"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

// 查询英雄抽奖信息
func handleHeroLottery(args []interface{}) {
	log.Debug("game handleHeroLottery")
	a := args[1].(gate.Agent)

	// // 输出收到的消息的内容
	log.Debug("user %v", a.UserData())
	player := a.UserData().(*entry.Player)

	response := new(msg.HeroLotteryResponse)
	response.Code = msg.ResponseCode_SUCCESS

	heroLottery := new(msg.HeroLottery)
	heroLottery.FreeGoodLottery = player.ExtendInfo.FreeGoodLottery
	heroLottery.FreeBetterLottery = player.ExtendInfo.FreeBetterLottery
	heroLottery.MaxFreeGoodLottery = player.ExtendInfo.MaxFreeGoodLottery
	heroLottery.MaxFreeBetterLottery = player.ExtendInfo.MaxFreeBetterLottery
	heroLottery.NextGoodLotteryStamp = player.ExtendInfo.LastFreeGoodLotteryStamp + 5*tool.SecondsPerMinute
	heroLottery.NextBetterLotteryStamp = player.ExtendInfo.LastFreeBetterLotteryStamp + tool.SecondsPerDay
	heroLottery.GoodLotteryCnt = player.ExtendInfo.GoodLotteryCnt
	heroLottery.BetterLotteryCnt = player.ExtendInfo.BetterLotteryCnt
	heroLottery.NeedGoodLotteryCnt = player.ExtendInfo.NeedGoodLotteryCnt
	heroLottery.NeedBetterLotteryCnt = player.ExtendInfo.NeedBetterLotteryCnt

	response.HeroLottery = heroLottery

	a.WriteMsg(response)
}

// 查询自己拥有的英雄
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

// 查询一个英雄的技能
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

func handleHeroEquips(args []interface{}) {
	log.Debug("game handleHeroEquips")

}

// 随机英雄
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
		if (player.ExtendInfo.FreeGoodLottery == 0 || player.ExtendInfo.LastFreeGoodLotteryStamp+5*tool.SecondsPerMinute > time.Now().Unix()) && player.BaseInfo.Gold < 1000 {
			response.Code = msg.ResponseCode_FAIL
			err := new(msg.Error)
			err.Code = define.HeroRandomGoldErr
			err.Msg = define.ERRMAP[err.Code]
			response.Err = err
		} else {
			rd := rand.Intn(len(heros))
			hero := entry.NewHero()
			tool.DeepCopy(hero, heros[rd])
			hero.SetHeroId(tool.UniqueId())
			hero.SetPlayerId(player.UserId)
			hero.SetExp(0)
			hero.SetLevel(1)
			hero.SetLevelUpExp(90)
			hero.SetSkillPoint(1)
			hero.SetMaxBlood(hero.Blood)
			hero.SetMaxMP(hero.MP)

			data.Module.InitHeroSkills(hero)
			data.Module.SavePlayerHero(player, hero)

			if player.ExtendInfo.FreeGoodLottery > 0 && player.ExtendInfo.LastFreeGoodLotteryStamp+5*tool.SecondsPerMinute < time.Now().Unix() {
				log.Debug("player.ExtendInfo.FreeGoodLottery === %v", player.ExtendInfo.FreeGoodLottery)
				player.ExtendInfo.SetFreeGoodLottery(player.ExtendInfo.FreeGoodLottery - 1)
				player.ExtendInfo.LastFreeGoodLotteryStamp = time.Now().Unix()
			} else {
				log.Debug("player.BaseInfo.Gold === %v", player.BaseInfo.Gold)
				player.BaseInfo.SetGold(player.BaseInfo.Gold - 1000)
			}

			player.ExtendInfo.SetGoodLotteryCnt(player.ExtendInfo.GoodLotteryCnt + 1)

			response.Hero = ConverHeroToMsgHero(hero)
			response.Code = msg.ResponseCode_SUCCESS
		}
		break
	case msg.HeroRandomLevel_BETTER:
		if (player.ExtendInfo.FreeBetterLottery == 0 || player.ExtendInfo.LastFreeBetterLotteryStamp+tool.SecondsPerDay > time.Now().Unix()) && player.BaseInfo.Diamond < 200 {
			response.Code = msg.ResponseCode_FAIL
			err := new(msg.Error)
			err.Code = define.HeroRandomDiamondErr
			err.Msg = define.ERRMAP[err.Code]
			response.Err = err
		} else {
			rd := rand.Intn(len(heros))
			hero := new(entry.Hero)
			tool.DeepCopy(hero, heros[rd])
			hero.SetHeroId(tool.UniqueId())
			hero.SetPlayerId(player.UserId)
			hero.SetExp(0)
			hero.SetLevel(1)
			hero.SetLevelUpExp(90)
			hero.SetSkillPoint(1)
			hero.SetMaxBlood(hero.Blood)
			hero.SetMaxMP(hero.MP)

			data.Module.SavePlayerHero(player, hero)

			if player.ExtendInfo.FreeBetterLottery > 0 && player.ExtendInfo.LastFreeBetterLotteryStamp+tool.SecondsPerDay < time.Now().Unix() {
				player.ExtendInfo.SetFreeBetterLottery(player.ExtendInfo.FreeBetterLottery - 1)
				player.ExtendInfo.LastFreeBetterLotteryStamp = time.Now().Unix()
			} else {
				player.BaseInfo.SetDiamond(player.BaseInfo.Diamond - 200)
			}

			player.ExtendInfo.SetBetterLotteryCnt(player.ExtendInfo.BetterLotteryCnt + 1)

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
			player.BaseInfo.SetDiamond(player.BaseInfo.Diamond - 200)

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

// 查询所有英雄
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

// 上阵一个英雄
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

// 下阵一个英雄
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

// TODO 英雄穿装备

// TODO 英雄卸下装备



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
	hero.SkillPoint = v.SkillPoint
	return hero
}
