package internal

import (
	"reflect"
	"server/data"
	"server/data/entry"
	"server/define"
	"server/msg"
	"time"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	handleMsg(&msg.RegisteRequest{}, handleRegiste)
	handleMsg(&msg.LoginRequest{}, handleAuth)
	handleMsg(&msg.ZoneRequest{}, handleAllZone)
}

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleRegiste(args []interface{}) {
	m := args[0].(*msg.RegisteRequest)
	a := args[1].(gate.Agent)

	log.Debug("[login handleRegiste] accountd = " + m.GetAccount() + " password = " + m.GetPassword())

	// 账号是否存在
	if data.Module.IsAccountExist(m.GetAccount()) {
		log.Debug("account exist ", m.GetAccount())
		response := new(msg.RegisteResponse)
		response.Code = msg.ResponseCode_FAIL
		response.Err = NewErr(define.LoginRegisteAccountExist)
		a.WriteMsg(response)
		return
	}

	// player := data.Module.FindPlayer(m.GetAccount(), m.GetPassword())

	// if player != nil {
	// 	log.Debug("user exist ", player.UserId, player.Name)
	// 	response := new(msg.RegisteResponse)
	// 	response.Code = msg.ResponseCode_FAIL
	// 	response.Err = NewErr(define.LoginRegisteExistErr)
	// 	a.WriteMsg(response)
	// 	return
	// }

	player := entry.NewPlayer()
	player.SetAccount(m.GetAccount())
	player.SetPassword(m.GetPassword())
	player.SetName(m.GetAccount())

	player.BaseInfo = entry.NewBaseInfo()
	player.ExtendInfo = entry.NewExtendInfo()

	data.Module.SavePlayer(player)

	response := new(msg.RegisteResponse)
	response.Code = msg.ResponseCode_SUCCESS
	response.Player = ConverPlayerToMsgPlayer(player)
	a.WriteMsg(response)
}

func handleAuth(args []interface{}) {
	m := args[0].(*msg.LoginRequest)
	a := args[1].(gate.Agent)

	log.Debug("[login handleAuth] accountd = " + m.GetAccount() + " password = " + m.GetPassword())

	player := data.Module.FindPlayer(m.GetAccount(), m.GetPassword())

	response := new(msg.LoginResponse)
	if player == nil {
		response.Code = msg.ResponseCode_FAIL
		response.Err = NewErr(define.LoginLoginNotExistErr)
		a.WriteMsg(response)
		return
	}

	player.SetLoginTime(time.Now())

	log.Debug("user exist ", player.UserId, player.Name)

	response.Code = msg.ResponseCode_SUCCESS
	response.Player = ConverPlayerToMsgPlayer(player)

	a.SetUserData(player)
	a.WriteMsg(response)
}

func handleAllZone(args []interface{}) {
	log.Debug("[login handleAllZone] ")

	a := args[1].(gate.Agent)

	response := new(msg.ZoneResponse)

	response.Code = msg.ResponseCode_SUCCESS
	response.Zones = make([]*msg.Zone, 0)
	for _, z := range data.Module.AllZones() {
		response.Zones = append(response.Zones, ConverZoneToMsgZone(z))
	}

	a.WriteMsg(response)
}

func ConverPlayerToMsgPlayer(v *entry.Player) *msg.Player {
	player := new(msg.Player)
	player.UserId = v.UserId
	player.Name = v.Name
	if v.BaseInfo != nil {
		player.BaseInfo = ConverBaseInfoToMsgBaseInfo(v.BaseInfo)
	}
	return player
}

func ConverBaseInfoToMsgBaseInfo(v *entry.BaseInfo) *msg.BaseInfo {
	baseInfo := new(msg.BaseInfo)
	baseInfo.Level = v.Level
	baseInfo.Gold = v.Gold
	baseInfo.Diamond = v.Diamond
	baseInfo.Exp = v.Exp
	baseInfo.Power = v.Power
	baseInfo.LevelUpExp = v.LevelUpExp
	baseInfo.MaxPower = v.MaxPower
	return baseInfo
}

func ConverZoneToMsgZone(v *entry.Zone) *msg.Zone {
	zone := new(msg.Zone)
	zone.Id = v.Id
	zone.Name = v.Name
	zone.IsNew = v.IsNew
	return zone
}

func NewErr(errCode int32) *msg.Error {
	err := new(msg.Error)
	err.Code = errCode
	err.Msg = define.ERRMAP[err.Code]
	return err
}
