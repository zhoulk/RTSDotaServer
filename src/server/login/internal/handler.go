package internal

import (
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
	handleMsg(&msg.RegisteRequest{}, handleRegiste)
	handleMsg(&msg.LoginRequest{}, handleAuth)
}

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleRegiste(args []interface{}) {
	m := args[0].(*msg.RegisteRequest)
	a := args[1].(gate.Agent)

	log.Debug("[login handleRegiste] accountd = " + m.GetAccount() + " password = " + m.GetPassword())

	player := data.Module.FindPlayer(m.GetAccount(), m.GetPassword())

	response := new(msg.RegisteResponse)

	if player != nil {
		log.Debug("user exist ", player.UserId, player.Name)
		response.Code = msg.ResponseCode_FAIL
		response.Err = NewErr(define.LoginRegisteExistErr)
		a.WriteMsg(response)
		return
	}

	player = new(entry.Player)
	player.UserId = tool.UniqueId()
	player.Account = m.GetAccount()
	player.Password = m.GetPassword()
	player.Name = m.GetAccount()

	baseInfo := new(entry.BaseInfo)
	baseInfo.Gold = 10000
	baseInfo.Diamond = 1000
	baseInfo.Level = 1
	baseInfo.Power = 120
	player.BaseInfo = baseInfo

	data.Module.SavePlayer(player)

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
	}

	log.Debug("user exist ", player.UserId, player.Name)

	response.Code = msg.ResponseCode_SUCCESS
	response.Player = ConverPlayerToMsgPlayer(player)

	a.SetUserData(player)
	a.WriteMsg(response)
}

func ConverPlayerToMsgPlayer(v *entry.Player) *msg.Player {
	player := new(msg.Player)
	player.UserId = v.UserId
	player.Name = v.Name
	player.BaseInfo = ConverBaseInfoToMsgBaseInfo(v.BaseInfo)
	return player
}

func ConverBaseInfoToMsgBaseInfo(v *entry.BaseInfo) *msg.BaseInfo {
	baseInfo := new(msg.BaseInfo)
	baseInfo.Level = v.Level
	baseInfo.Gold = v.Gold
	baseInfo.Diamond = v.Diamond
	baseInfo.Exp = v.Exp
	baseInfo.Power = v.Power
	return baseInfo
}

func NewErr(errCode int32) *msg.Error {
	err := new(msg.Error)
	err.Code = errCode
	err.Msg = define.ERRMAP[err.Code]
	return err
}
