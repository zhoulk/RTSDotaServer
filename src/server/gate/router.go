package gate

import (
	"server/login"
	"github.com/name5566/leaf/log"

	"server/msg"
)

func init()  {
	log.Debug("gate init")

	msg.Processor.SetRouter(&msg.LoginRequest{}, login.ChanRPC)
}