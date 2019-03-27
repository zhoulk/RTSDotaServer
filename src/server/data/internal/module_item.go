package internal

import "server/data/entry"

func (m *Module) AllItems() []*entry.Item {
	return m.items
}
