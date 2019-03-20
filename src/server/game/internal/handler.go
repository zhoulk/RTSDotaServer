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
		hero := new(msg.Hero)
		hero.Id = v.Id
		hero.Name = v.Name
		hero.Level = v.Level
		hero.Strength = v.Strength
		hero.Agility = v.Agility
		hero.Intelligence = v.Intelligence
		hero.Armor = v.Armor
		hero.Attack = v.Attack
		hero.Blood = v.Blood
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
		hero := new(msg.Hero)
		hero.Id = v.Id
		hero.Name = v.Name
		hero.Level = v.Level
		hero.Strength = v.Strength
		hero.Agility = v.Agility
		hero.Intelligence = v.Intelligence
		hero.Armor = v.Armor
		hero.Attack = v.Attack
		hero.Blood = v.Blood
		heros = append(heros, hero)
	}
	response.Heros = heros

	a.WriteMsg(response)
}
