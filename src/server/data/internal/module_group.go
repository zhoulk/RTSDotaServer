package internal

import "server/data/entry"

func (m *Module) OwnGroup(player *entry.Player) *entry.Group {
	if player == nil || len(player.UserId) == 0 {
		return nil
	}

	if player.ExtendInfo == nil || len(player.ExtendInfo.GroupId) == 0 {
		return nil
	}

	if m.groups != nil {
		for _, group := range m.groups {
			if group.GroupId == player.ExtendInfo.GroupId {
				return group
			}
		}
	}

	return nil
}
