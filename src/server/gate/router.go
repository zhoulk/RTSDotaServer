package gate

import (
	"server/game"
	"server/login"
	"server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.LoginRequest{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.RegisteRequest{}, login.ChanRPC)

	msg.Processor.SetRouter(&msg.HeroRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.HeroRandomRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.HeroOwnRequest{}, game.ChanRPC)
}
