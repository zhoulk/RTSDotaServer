package internal

import (
	"errors"
	"server/data/entry"
)

func (m *Module) CalPlayerPower() {
	for _, player := range m.players {
		// if player.BaseInfo.Power < player.BaseInfo.MaxPower {
		player.BaseInfo.SetPower(player.BaseInfo.Power + 1)
		// }
	}
}

func (m *Module) AllZones() []*entry.Zone {
	return m.zones
}

func (m *Module) AllExpHeros() []int32 {
	return m.heroExpList
}

func (m *Module) AllExpPlayers() []int32 {
	return m.playerExpList
}

func (m *Module) AllPlayers() map[string]*entry.Player {
	return m.players
}

func (m *Module) FindPlayer(account string, pwd string) *entry.Player {
	player := m.players[account]
	if player != nil && player.Password == pwd {

	} else {
		for _, p := range m.players {
			if p.UserId == account {
				player = p
				break
			}
		}
	}
	return player
}

func (m *Module) IsAccountExist(account string) bool {
	player := m.players[account]
	return player != nil
}

func (m *Module) SavePlayer(player *entry.Player) error {
	if player == nil || len(player.UserId) == 0 {
		return errors.New("player is nil or userId length is 0")
	}
	m.players[player.Account] = player
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

func (m *Module) EffectPlayerByEarn(p *entry.Player, earn *entry.Earn) {
	p.BaseInfo.Exp += earn.PlayerExp
	for {
		if p.BaseInfo.Exp < p.BaseInfo.LevelUpExp {
			break
		}
		if p.BaseInfo.Level+1 >= int32(len(m.playerExpList)) {
			break
		}
		m.PlayerLevelUp(p)
	}

	p.IsDirty = true
}

func (m *Module) PlayerLevelUp(p *entry.Player) {
	p.BaseInfo.Exp -= p.BaseInfo.LevelUpExp
	p.BaseInfo.Level += 1
	p.BaseInfo.LevelUpExp = m.playerExpList[p.BaseInfo.Level]
}
