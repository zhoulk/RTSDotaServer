package internal

import (
	"server/data"
	"server/data/entry"
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

	a.WriteMsg(response)
}

func handleGroupList(args []interface{}) {
	log.Debug("game handleGroupList")
}

func handleGroupCreate(args []interface{}) {
	log.Debug("game handleGroupCreate")
}

func handleGroupMembers(args []interface{}) {
	log.Debug("game handleGroupMembers")
}

func ConverGroupToMsgGroup(v *entry.Group) *msg.Group {
	group := new(msg.Group)
	group.GroupId = v.GroupId
	group.GroupName = v.GroupName
	group.GroupDeclaration = v.GroupDeclaration
	group.GroupLeader = v.GroupLeader
	group.MemberCnt = v.MemberCnt
	group.MemberTotal = v.MemberTotal
	return group
}
