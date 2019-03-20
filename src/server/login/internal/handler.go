package internal

import (
	"reflect"
	"server/data"
	"server/data/entry"
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
	} else {
		player = new(entry.Player)
		player.UserId = tool.UniqueId()
		player.Account = m.GetAccount()
		player.Password = m.GetPassword()
		player.Name = m.GetAccount()

		baseInfo := new(entry.BaseInfo)
		baseInfo.Gold = 10000
		baseInfo.Diamond = 1000
		player.BaseInfo = baseInfo

		data.Module.SavePlayer(player)

		response.Code = msg.ResponseCode_SUCCESS
		response.Uid = player.UserId
	}

	a.WriteMsg(response)
}

func handleAuth(args []interface{}) {
	m := args[0].(*msg.LoginRequest)
	a := args[1].(gate.Agent)

	log.Debug("[login handleAuth] accountd = " + m.GetAccount() + " password = " + m.GetPassword())

	player := data.Module.FindPlayer(m.GetAccount(), m.GetPassword())

	response := new(msg.LoginResponse)
	if player != nil {
		log.Debug("user exist ", player.UserId, player.Name)

		response.Code = msg.ResponseCode_SUCCESS
		response.Uid = player.UserId

		a.SetUserData(player)
	} else {
		response.Code = msg.ResponseCode_FAIL
	}

	a.WriteMsg(response)
}
