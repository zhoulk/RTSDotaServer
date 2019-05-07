package internal

import (
	"server/data/entry"
	"server/tool"
)

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

func (m *Module) AllGroups() []*entry.Group {
	return m.groups
}

func (m *Module) FindGroup(groupId string) *entry.Group {
	for _, gp := range m.groups {
		if gp.GroupId == groupId {
			return gp
		}
	}
	return nil
}

func (m *Module) IsInGroup(player *entry.Player, groupId string) bool {
	if player == nil || len(player.UserId) == 0 {
		return false
	}

	mems := m.GroupMembers(groupId)
	if mems != nil {
		for _, mem := range mems {
			if mem.UserId == player.UserId {
				return true
			}
		}
	}

	return false
}

func (m *Module) CreateGroup(player *entry.Player, name string) *entry.Group {
	if player == nil || len(player.UserId) == 0 {
		return nil
	}

	group := entry.NewGroup()
	group.SetGroupId(tool.UniqueId())
	group.SetGroupName(name)
	group.SetGroupLeader(player.Name)
	group.SetMemberCnt(1)
	group.SetMemberTotal(30)
	group.SetGroupLevel(1)
	group.SetContriCurrent(0)
	group.SetContriLevelUp(5000)
	group.GroupMembers = make([]*entry.GroupMember, 0)

	member := entry.NewGroupMember()
	member.SetUserId(player.UserId)
	member.SetLevel(player.BaseInfo.Level)
	member.SetName(player.Name)
	group.GroupMembers = append(group.GroupMembers, member)

	if player.ExtendInfo == nil {
		player.ExtendInfo = entry.NewExtendInfo()
	}
	player.ExtendInfo.SetGroupId(group.GroupId)
	player.BaseInfo.SetGold(player.BaseInfo.Gold - m.gameConfig.GroupPrice)

	m.groups = append(m.groups, group)

	return group
}

func (m *Module) GroupMembers(groupId string) []*entry.GroupMember {
	for _, group := range m.groups {
		if group.GroupId == groupId {
			return group.GroupMembers
		}
	}
	return nil
}
