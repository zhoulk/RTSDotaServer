package internal

import (
	"server/data/entry"
	"time"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	log.Debug("[login] create agent")
	a := args[0].(gate.Agent)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	log.Debug("[login] close agent")
	a := args[0].(gate.Agent)

	player := a.UserData().(*entry.Player)

	player.SetLogoutTime(time.Now())

	_ = a
}
