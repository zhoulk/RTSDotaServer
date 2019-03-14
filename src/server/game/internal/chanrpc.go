package internal

import (
	"github.com/name5566/leaf/log"

	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	log.Debug("[game] create agent")
	a := args[0].(gate.Agent)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	log.Debug("[game] close agent")
	a := args[0].(gate.Agent)
	_ = a
}
