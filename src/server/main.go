package main

import (
	"server/data"
	"server/game"
	"server/gate"
	"server/login"

	"github.com/name5566/leaf"
)

func main() {

	// var t string
	// flag.StringVar(&t, "t", "", "server type")
	// var num int
	// flag.IntVar(&num, "num", 0, "server num")

	// flag.Parse()
	// log.Debug("[Start] %v[%v]", t, num)

	// lconf.LogLevel = conf.Server.LogLevel
	// lconf.LogPath = conf.Server.LogPath
	// lconf.LogFlag = conf.LogFlag
	// lconf.ConsolePort = conf.Server.ConsolePort
	// lconf.ProfilePath = conf.Server.ProfilePath

	// leaf.Run(
	// 	gate.Module,
	// 	login.Module,
	// 	game.Module,
	// )

	leaf.Run(
		gate.Module,
		game.Module,
		login.Module,
		data.Module,
	)
}
