package internal

import "server/data/entry"

func (m *Module) AddGuanKaBattle(battleId string, gk *entry.GuanKa) {
	info := new(entry.BattleInfo)
	info.BattleId = battleId
	info.Type = entry.BattleTypeGuanKa
	info.Guanka = gk
	m.battleCache[battleId] = info
}

func (m *Module) FindBattle(battleId string) *entry.BattleInfo {
	return m.battleCache[battleId]
}
