package internal

import (
	"reflect"
	"server/msg"

	"github.com/name5566/leaf/gate"

	"github.com/name5566/leaf/log"
)

func init() {
	log.Debug("login init")
	handleMsg(&msg.LoginRequest{}, handleAuth)
}

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleAuth(args []interface{}) {
	m := args[0].(*msg.LoginRequest)
	a := args[1].(gate.Agent)

	log.Debug("[login handleAuth] accountd = " + m.GetAccount() + " password = " + m.GetPassword())

	a.WriteMsg(&msg.LoginResponse{
		Code: msg.ResponseCode_SUCCESS,
		Uid:  "1234567890",
	})
}
