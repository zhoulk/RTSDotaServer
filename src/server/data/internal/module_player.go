package internal

import (
	"errors"
	"server/data/entry"
)

func (m *Module) FindPlayer(account string, pwd string) *entry.Player {
	player := m.players[account+"-"+pwd]
	if player == nil {
		for _, p := range m.players {
			if p.UserId == account {
				player = p
				break
			}
		}
	}
	return player
}

func (m *Module) SavePlayer(player *entry.Player) error {
	if player == nil || len(player.UserId) == 0 {
		return errors.New("player is nil or userId length is 0")
	}
	m.players[player.Account+"-"+player.Password] = player
	return nil
}

func (m *Module) SavePlayerHero(player *entry.Player, hero *entry.Hero) error {
	if player == nil || len(player.UserId) == 0 {
		return errors.New("player is nil or userId length is 0")
	}
	if hero == nil {
		return errors.New("hero is nil")
	}

	heros := m.playerHeros[player.UserId]
	if heros == nil {
		heros = make([]*entry.Hero, 0)
	}
	if hero.Skills == nil {
		hero.Skills = m.FindHeroSkills(hero)
	}
	heros = append(heros, hero)
	m.playerHeros[player.UserId] = heros

	return nil
}
