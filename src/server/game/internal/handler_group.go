package internal

import (
	"server/data"
	"server/data/entry"
	"server/define"
	"server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

// 获取自己参加的军团
func handleGroupOwn(args []interface{}) {
	log.Debug("game handleGroupOwn")

	a := args[1].(gate.Agent)

	// // 输出收到的消息的内容
	log.Debug("user %v", a.UserData())
	player := a.UserData().(*entry.Player)

	response := new(msg.GroupOwnResponse)
	response.Code = msg.ResponseCode_SUCCESS

	ownGroup := data.Module.OwnGroup(player)
	if ownGroup != nil {
		response.Group = ConverGroupToMsgGroup(ownGroup)
	}

	log.Debug("%v", response.Group)

	a.WriteMsg(response)
}

// 获取军团列表
func handleGroupList(args []interface{}) {
	log.Debug("game handleGroupList")

	a := args[1].(gate.Agent)

	// // 输出收到的消息的内容
	log.Debug("user %v", a.UserData())

	response := new(msg.GroupListResponse)
	response.Code = msg.ResponseCode_SUCCESS

	groups := data.Module.AllGroups()
	response.Groups = make([]*msg.Group, 0)
	for _, group := range groups {
		response.Groups = append(response.Groups, ConverGroupToMsgGroup(group))
	}

	a.WriteMsg(response)
}

// 创建军团
func handleGroupCreate(args []interface{}) {
	log.Debug("game handleGroupCreate")

	m := args[0].(*msg.GroupCreateRequest)
	a := args[1].(gate.Agent)

	// // 输出收到的消息的内容
	log.Debug("user %v", a.UserData())
	groupName := m.GetGroupName()
	player := a.UserData().(*entry.Player)

	if player.BaseInfo.Diamond < data.Module.GameConfig().GroupPrice {
		response := new(msg.GroupCreateResponse)
		response.Code = msg.ResponseCode_FAIL
		response.Err = NewErr(define.GroupCreateDiamondErr)
		a.WriteMsg(response)
		return
	}

	response := new(msg.GroupCreateResponse)
	response.Code = msg.ResponseCode_SUCCESS

	ownGroup := data.Module.CreateGroup(player, groupName)
	if ownGroup != nil {
		response.Group = ConverGroupToMsgGroup(ownGroup)
	}

	a.WriteMsg(response)
}

// 获取军团成员
func handleGroupMembers(args []interface{}) {
	log.Debug("game handleGroupMembers")

	m := args[0].(*msg.GroupMembersRequest)
	a := args[1].(gate.Agent)

	// // 输出收到的消息的内容
	log.Debug("user %v", a.UserData())
	groupID := m.GetGroupId()
	// player := a.UserData().(*entry.Player)

	response := new(msg.GroupMembersResponse)
	response.Code = msg.ResponseCode_SUCCESS

	members := data.Module.GroupMembers(groupID)
	response.Members = make([]*msg.GroupMember, 0)
	if members != nil {
		for _, member := range members {
			response.Members = append(response.Members, ConvertGroupMemberToMsgGroupMember(member))
		}
	}

	a.WriteMsg(response)
}

// TODO 申请加入
func handleGroupApply(args []interface{}) {
	log.Debug("game handleGroupApply")

	m := args[0].(*msg.GroupApplyRequest)
	a := args[1].(gate.Agent)

	// // 输出收到的消息的内容
	log.Debug("user %v", a.UserData())
	player := a.UserData().(*entry.Player)

	groupId := m.GroupId

	gp := data.Module.FindGroup(groupId)
	if gp == nil {
		response := new(msg.GroupApplyResponse)
		response.Code = msg.ResponseCode_FAIL
		response.Err = NewErr(define.GroupApplyExistErr)
		a.WriteMsg(response)
		return
	}

	res := data.Module.IsInGroup(player, groupId)
	if res {
		response := new(msg.GroupApplyResponse)
		response.Code = msg.ResponseCode_FAIL
		response.Err = NewErr(define.GroupApplyIsInErr)
		a.WriteMsg(response)
		return
	}

	data.Module.GroupApply(player, groupId)

	response := new(msg.GroupApplyResponse)
	response.Code = msg.ResponseCode_SUCCESS
	a.WriteMsg(response)
}

// TODO 获取申请成员
func handleGroupApplyMembers(args []interface{}) {
	log.Debug("game handleGroupApply")

	// m := args[0].(*msg.GroupMembersRequest)
	// a := args[1].(gate.Agent)

	// // // 输出收到的消息的内容
	// log.Debug("user %v", a.UserData())
}

// TODO 通过、拒绝、剔除 成员
func handleGroupOper(args []interface{}) {
	log.Debug("game handleGroupApply")

	// m := args[0].(*msg.GroupMembersRequest)
	// a := args[1].(gate.Agent)

	// // // 输出收到的消息的内容
	// log.Debug("user %v", a.UserData())
}

// TODO 捐献贡献
func handleGroupContri(args []interface{}) {
	log.Debug("game handleGroupApply")

	// m := args[0].(*msg.GroupMembersRequest)
	// a := args[1].(gate.Agent)

	// // // 输出收到的消息的内容
	// log.Debug("user %v", a.UserData())
}

// TODO 升级军团科技

// ConverGroupToMsgGroup ...
func ConverGroupToMsgGroup(v *entry.Group) *msg.Group {
	group := new(msg.Group)
	group.GroupId = v.GroupId
	group.GroupName = v.GroupName
	group.GroupDeclaration = v.GroupDeclaration
	group.GroupLeader = v.GroupLeader
	group.MemberCnt = v.MemberCnt
	group.MemberTotal = v.MemberTotal
	group.GroupLevel = v.GroupLevel
	group.ContriCurrent = v.ContriCurrent
	group.ContriLevelUp = v.ContriLevelUp
	return group
}

// ConvertGroupMemberToMsgGroupMember ...
func ConvertGroupMemberToMsgGroupMember(v *entry.GroupMember) *msg.GroupMember {
	member := new(msg.GroupMember)
	member.UserId = v.UserId
	member.Level = v.Level
	member.Name = v.Name
	member.OffLineTime = v.OffLineTime
	member.Power = v.Power
	member.ContriToday = v.ContriToday
	member.ContriTotal = v.ContriTotal
	return member
}
