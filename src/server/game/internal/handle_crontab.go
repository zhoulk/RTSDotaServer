package internal

import (
	"server/data"
	"server/data/entry"
	"server/msg"
	"time"
)

func handleCrontab() {
	go executePersistent()
	go executeOneMinite()
}

func executePersistent() {
	timer := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-timer.C:
			data.Module.DoPersistent()
		}
	}
}

func executeOneMinite() {
	timer := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-timer.C:
			data.Module.CalPlayerPower()
			notifyPlayerInfo()
		}
	}
}

func notifyPlayerInfo() {
	for _, player := range data.Module.AllPlayers() {
		if player.BaseInfo != nil && player.BaseInfo.IsNotify {
			msg := new(msg.PlayerInfoNotify)
			msg.Player = entry.ConverPlayerToMsgPlayer(player)
			if player.Agent != nil {
				player.Agent.WriteMsg(msg)
				player.BaseInfo.IsNotify = false
			}
		}
	}
}
