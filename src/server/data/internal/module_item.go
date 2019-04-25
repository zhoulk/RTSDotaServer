package internal

import "server/data/entry"

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
