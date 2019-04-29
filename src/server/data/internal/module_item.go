package internal

import (
	"server/data/entry"
)

func (m *Module) HasItem(id int32) bool {
	exist := false
	for _, define := range m.items {
		if define.Id == id {
			exist = true
			break
		}
	}

	return exist
}

func (m *Module) AllItems() []*entry.Item {
	return m.items
}

func (m *Module) AllOwnEquips(p *entry.Player) []*entry.Item {
	return m.allOwnItems(p, entry.ItemTypeEquip)
}

func (m *Module) AllOwnConsumes(p *entry.Player) []*entry.Item {
	return m.allOwnItems(p, entry.ItemTypeConsume)
}

func (m *Module) AllOwnHeroChips(p *entry.Player) []*entry.Item {
	return m.allOwnItems(p, entry.ItemTypeHeroChip)
}

func (m *Module) allOwnItems(p *entry.Player, itemType int32) []*entry.Item {
	if p == nil || len(p.UserId) == 0 {
		return nil
	}
	items := m.playerItems[p.UserId]
	if items == nil {
		return nil
	} else {
		res := make([]*entry.Item, 0)
		for _, item := range items {
			if itemType == item.Type {
				res = append(res, item)
			}
		}
		return res
	}
}

func (m *Module) EffectItemsByEarn(p *entry.Player, earn *entry.Earn) {
	if p == nil || len(p.UserId) == 0 {
		return
	}
	items := m.playerItems[p.UserId]
	if items == nil {
		items = make([]*entry.Item, 0)
	}
	for _, item := range earn.Items {
		items = append(items, item)
	}
	m.playerItems[p.UserId] = items
}
