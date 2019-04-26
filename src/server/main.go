package main

import (
	"flag"
	"server/conf"
	"server/data"
	"server/data/entry"
	"server/game"
	"server/gate"
	"server/login"

	"github.com/name5566/leaf"
	"github.com/name5566/leaf/log"
)

func main() {

	var serverId string
	flag.StringVar(&serverId, "s", "", "server id")
	// var num int
	// flag.IntVar(&num, "num", 0, "server num")

	flag.Parse()
	var zone *entry.Zone
	for _, z := range entry.ZoneList {
		if z.Id == serverId {
			zone = z
			break
		}
	}
	if zone == nil {
		log.Fatal("[Start ] serverId is invalid !  serverId = %v", serverId)
	}

	log.Debug("[Start] %v[%v]", serverId, zone)

	conf.Server.MaxConnNum = int(zone.MaxConnNum)
	conf.Server.TCPAddr = zone.TCPAddr

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

	// arr := tool.C_M_N(8, 3)
	// log.Debug("%v", arr)

	leaf.Run(
		gate.Module,
		game.Module,
		login.Module,
		data.Module,
	)
}
