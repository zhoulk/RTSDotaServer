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

func (m *Module) CreateGroup(player *entry.Player, name string) *entry.Group {
	if player == nil || len(player.UserId) == 0 {
		return nil
	}

	group := new(entry.Group)
	group.GroupId = tool.UniqueId()
	group.GroupName = name
	group.GroupLeader = player.Name
	group.MemberCnt = 1
	group.MemberTotal = 30
	group.GroupLevel = 1
	group.ContriCurrent = 0
	group.ContriLevelUp = 5000
	group.GroupMembers = make([]*entry.GroupMember, 0)

	member := new(entry.GroupMember)
	member.UserId = player.UserId
	member.Level = player.BaseInfo.Level
	member.Name = player.Name
	group.GroupMembers = append(group.GroupMembers, member)

	if player.ExtendInfo == nil {
		player.ExtendInfo = new(entry.ExtendInfo)
	}
	player.ExtendInfo.GroupId = group.GroupId

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
