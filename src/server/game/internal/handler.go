package internal

import (
	"math/rand"
	"reflect"
	"server/data"
	"server/data/entry"
	"server/define"
	"server/msg"
	"server/tool"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
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
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

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
	ownHeros, _ := data.Module.AllOwnHeros(player)
	for _, v := range ownHeros {
		hero := ConverHeroToMsgHero(v)
		heros = append(heros, hero)
	}
	response.Heros = heros

	a.WriteMsg(response)
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

func handleAllSkill(args []interface{}) {
	log.Debug("game handleAllSkill")

	//m := args[0].(*msg.HeroRequest)
	a := args[1].(gate.Agent)

	// // 输出收到的消息的内容
	log.Debug("user %v", a.UserData())

	response := new(msg.SkillResponse)
	response.Code = msg.ResponseCode_SUCCESS

	skills := make([]*msg.Skill, 0)
	for _, v := range data.Module.AllSkills() {
		skill := ConverSkillToMsgSkill(v)
		skills = append(skills, skill)
	}
	response.Skills = skills

	a.WriteMsg(response)
}

func handleAllItem(args []interface{}) {
	log.Debug("game handleAllItem")

	//m := args[0].(*msg.HeroRequest)
	a := args[1].(gate.Agent)

	// // 输出收到的消息的内容
	log.Debug("user %v", a.UserData())

	response := new(msg.ItemResponse)
	response.Code = msg.ResponseCode_SUCCESS

	items := make([]*msg.Item, 0)
	for _, v := range data.Module.AllItems() {
		item := ConverItemToMsgItem(v)
		items = append(items, item)
	}
	response.Items = items

	a.WriteMsg(response)
}

func handleAllChapter(args []interface{}) {
	log.Debug("game handleAllChapter")

	//m := args[0].(*msg.HeroRequest)
	a := args[1].(gate.Agent)

	player := a.UserData().(*entry.Player)

	response := new(msg.ChapterResponse)
	response.Code = msg.ResponseCode_SUCCESS

	chapters := make([]*msg.Chapter, 0)
	for _, v := range data.Module.AllChapters(player) {
		item := ConverChapterToMsgChapter(v)
		chapters = append(chapters, item)
	}
	response.Chapters = chapters

	a.WriteMsg(response)
}

func handleAllGuanKa(args []interface{}) {
	log.Debug("game handleAllGuanKa")

	//m := args[0].(*msg.HeroRequest)
	a := args[1].(gate.Agent)

	player := a.UserData().(*entry.Player)

	response := new(msg.GuanKaResponse)
	response.Code = msg.ResponseCode_SUCCESS

	guanKas := make([]*msg.GuanKa, 0)
	for _, v := range data.Module.AllGuanKas(player) {
		guanKa := ConverGuanKaToMsgGuanKa(v)
		guanKas = append(guanKas, guanKa)
	}
	response.GuanKas = guanKas

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

	response := new(msg.BattleGuanKaResponse)
	response.Code = msg.ResponseCode_SUCCESS

	response.Result = define.BattleResult_Success
	response.Guanka = ConverGuanKaToMsgGuanKa(guanKa)

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
	for _, id := range v.SkillIds {
		hero.SkillIds = append(hero.SkillIds, tool.String(id))
	}
	hero.HeroId = v.HeroId
	hero.PlayerId = v.PlayerId
	hero.IsSelect = v.IsSelect
	hero.Pos = v.Pos
	return hero
}

func ConverSkillToMsgSkill(v *entry.Skill) *msg.Skill {
	skill := new(msg.Skill)
	skill.Id = v.Id
	skill.Name = v.Name
	skill.Level = v.Level
	skill.Type = v.Type
	skill.Desc = v.Desc
	return skill
}

func ConverItemToMsgItem(v *entry.Item) *msg.Item {
	item := new(msg.Item)
	item.Id = v.Id
	item.Name = v.Name
	item.Price = v.Price
	item.Effect = v.Effect
	item.Desc = v.Desc
	item.Mixs = make([]*msg.Mix, 0)
	if v.Mixs != nil {
		for _, mix := range v.Mixs {
			item.Mixs = append(item.Mixs, ConverMixToMsgMix(mix))
		}
	}
	return item
}

func ConverMixToMsgMix(v *entry.Mix) *msg.Mix {
	mix := new(msg.Mix)
	mix.ItemId = v.ItemId
	mix.Num = v.Num
	return mix
}

func ConverChapterToMsgChapter(v *entry.Chapter) *msg.Chapter {
	chapter := new(msg.Chapter)
	chapter.Id = v.Id
	chapter.Name = v.Name
	chapter.IsOpen = v.IsOpen
	return chapter
}

func ConverGuanKaToMsgGuanKa(v *entry.GuanKa) *msg.GuanKa {
	guanKa := new(msg.GuanKa)
	guanKa.Id = v.Id
	guanKa.Name = v.Name
	guanKa.ChapterId = v.ChapterId
	guanKa.IsOpen = v.IsOpen
	guanKa.Earn = ConverEarnToMsgEarn(v.Earn)
	guanKa.Expend = ConverExpendToMsgExpend(v.Expend)
	return guanKa
}

func ConverEarnToMsgEarn(v *entry.Earn) *msg.Earn {
	eran := new(msg.Earn)
	for _, id := range v.ItemIds {
		eran.ItemIds = append(eran.ItemIds, tool.String(id))
	}
	eran.HeroExp = v.HeroExp
	eran.PlayerExp = v.PlayerExp
	eran.Gold = v.Gold
	return eran
}

func ConverExpendToMsgExpend(v *entry.Expend) *msg.Expend {
	expend := new(msg.Expend)
	expend.Power = v.Power
	return expend
}

func NewErr(errCode int32) *msg.Error {
	err := new(msg.Error)
	err.Code = errCode
	err.Msg = define.ERRMAP[err.Code]
	return err
}
