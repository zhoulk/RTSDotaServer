package internal

import (
	"server/data"
	"server/data/entry"
	"server/define"
	"server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

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

func handleGroupList(args []interface{}) {
	log.Debug("game handleGroupList")
}

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
