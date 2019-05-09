package internal

import (
	"server/data/entry"
	"server/define"
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

	return m.IsUserInGroup(player.UserId, groupId)
}

func (m *Module) IsUserInGroup(userId string, groupId string) bool {
	mems := m.GroupMembers(groupId)
	if mems != nil {
		for _, mem := range mems {
			if mem.UserId == userId {
				return true
			}
		}
	}

	return false
}

func (m *Module) GroupApply(player *entry.Player, groupId string) {
	if player == nil || len(player.UserId) == 0 {
		return
	}

	gp := m.FindGroup(groupId)
	if gp != nil {
		exist := false
		for _, mem := range gp.ApplyMembers {
			if mem.UserId == player.UserId {
				exist = true
			}
		}
		if !exist {
			mem := entry.NewGroupMember()
			mem.SetUserId(player.UserId)
			gp.ApplyMembers = append(gp.ApplyMembers, mem)
		}
	}
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
	group.ApplyMembers = make([]*entry.GroupMember, 0)

	member := entry.NewGroupMember()
	member.SetUserId(player.UserId)
	member.SetLevel(player.BaseInfo.Level)
	member.SetName(player.Name)
	member.SetJob(entry.GroupMemberJob_Leader)
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
	members := make([]*entry.GroupMember, 0)
	for _, group := range m.groups {
		if group.GroupId == groupId {
			for _, mem := range group.GroupMembers {
				if mem.LastOper == define.GroupOper_None {
					player := m.FindPlayer(mem.UserId, "")
					mem.Level = player.BaseInfo.Level
					mem.Name = player.Name
					mem.Power = player.BaseInfo.Military
					members = append(members, mem)
				}
			}
			break
		}
	}
	return members
}

func (m *Module) GroupApplyMembers(groupId string) []*entry.GroupMember {
	members := make([]*entry.GroupMember, 0)
	for _, group := range m.groups {
		if group.GroupId == groupId {
			for _, mem := range group.ApplyMembers {
				if mem.LastOper == define.GroupOper_None {
					player := m.FindPlayer(mem.UserId, "")
					mem.Level = player.BaseInfo.Level
					mem.Name = player.Name
					mem.Power = player.BaseInfo.Military
					members = append(members, mem)
				}
			}
			break
		}
	}
	return members
}

func (m *Module) GroupAgree(player *entry.Player, groupId string, userId string) int32 {
	if player == nil || len(player.UserId) == 0 {
		return 0
	}

	gp := m.FindGroup(groupId)
	if gp == nil {
		return define.GroupOperExistErr
	}
	isIn := m.IsUserInGroup(userId, groupId)
	if isIn {
		return define.GroupOperIsInErr
	}

	if gp.MemberCnt >= gp.MemberTotal {
		return define.GroupOperMemberFullErr
	}

	isLeader := false
	for _, mem := range gp.GroupMembers {
		if mem.UserId == player.UserId {
			if mem.Job == entry.GroupMemberJob_Leader || mem.Job == entry.GroupMemberJob_SecondLeader {
				isLeader = true
			}
		}
	}
	if !isLeader {
		return define.GroupOperNoLeaderErr
	}

	index := -1
	for i, mem := range gp.ApplyMembers {
		if mem.UserId == userId {
			mem.SetLastOper(define.GroupOper_Agree)
			index = i
		}
	}
	if index == -1 {
		return define.GroupOperNoApplyErr
	}

	member := entry.NewGroupMember()
	member.SetUserId(userId)
	member.SetJob(entry.GroupMemberJob_Member)
	gp.GroupMembers = append(gp.GroupMembers, member)

	p := m.FindPlayer(userId, "")
	p.ExtendInfo.SetGroupId(groupId)

	return 0
}

func (m *Module) GroupReject(player *entry.Player, groupId string, userId string) int32 {
	if player == nil || len(player.UserId) == 0 {
		return 0
	}

	gp := m.FindGroup(groupId)
	if gp == nil {
		return define.GroupOperExistErr
	}
	isIn := m.IsUserInGroup(userId, groupId)
	if isIn {
		return define.GroupOperIsInErr
	}

	isLeader := false
	for _, mem := range gp.GroupMembers {
		if mem.UserId == player.UserId {
			if mem.Job == entry.GroupMemberJob_Leader || mem.Job == entry.GroupMemberJob_SecondLeader {
				isLeader = true
			}
		}
	}
	if !isLeader {
		return define.GroupOperNoLeaderErr
	}

	index := -1
	for i, mem := range gp.ApplyMembers {
		if mem.UserId == userId {
			mem.SetLastOper(define.GroupOper_Reject)
			index = i
		}
	}
	if index == -1 {
		return define.GroupOperNoApplyErr
	}

	return 0
}

func (m *Module) GroupDel(player *entry.Player, groupId string, userId string) int32 {
	if player == nil || len(player.UserId) == 0 {
		return 0
	}

	gp := m.FindGroup(groupId)
	if gp == nil {
		return define.GroupOperExistErr
	}
	isIn := m.IsUserInGroup(userId, groupId)
	if !isIn {
		return define.GroupOperIsNotInErr
	}

	isLeader := false
	for _, mem := range gp.GroupMembers {
		if mem.UserId == player.UserId {
			if mem.Job == entry.GroupMemberJob_Leader || mem.Job == entry.GroupMemberJob_SecondLeader {
				isLeader = true
			}
		}
	}
	if !isLeader {
		return define.GroupOperNoLeaderErr
	}

	for _, mem := range gp.GroupMembers {
		if mem.UserId == userId {
			mem.SetLastOper(define.GroupOper_Del)
		}
	}

	return 0
}
