package internal

import (
	"errors"
	"server/base"
	"server/data/entry"

	"github.com/name5566/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
	players map[string]*entry.Player
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	m.players = make(map[string]*entry.Player)
}

func (m *Module) OnDestroy() {

}

func (m *Module) FindPlayer(userId string) *entry.Player {
	return m.players[userId]
}

func (m *Module) SavePlayer(player *entry.Player) error {
	if player == nil || len(player.UserId) == 0 {
		return errors.New("player is nil or userId length is 0")
	}
	m.players[player.UserId] = player
	return nil
}
