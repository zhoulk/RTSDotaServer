package entry

type Group struct {
	GroupId          string
	GroupName        string
	GroupLeader      string
	GroupDeclaration string
	MemberCnt        int32
	MemberTotal      int32
	GroupLevel       int32
	ContriCurrent    int32
	ContriLevelUp    int32

	GroupMembers []*GroupMember

	IsDirty bool
}

func NewGroup() *Group {
	group := new(Group)
	group.IsDirty = true
	return group
}

func (g *Group) SetGroupId(groupId string) {
	g.GroupId = groupId
	g.IsDirty = true
}

func (g *Group) SetGroupName(groupName string) {
	g.GroupName = groupName
	g.IsDirty = true
}

func (g *Group) SetGroupLeader(groupLeader string) {
	g.GroupLeader = groupLeader
	g.IsDirty = true
}

func (g *Group) SetGroupDeclaration(groupDeclaration string) {
	g.GroupDeclaration = groupDeclaration
	g.IsDirty = true
}

func (g *Group) SetMemberCnt(memberCnt int32) {
	g.MemberCnt = memberCnt
	g.IsDirty = true
}

func (g *Group) SetMemberTotal(memberTotal int32) {
	g.MemberTotal = memberTotal
	g.IsDirty = true
}

func (g *Group) SetGroupLevel(groupLevel int32) {
	g.GroupLevel = groupLevel
	g.IsDirty = true
}

func (g *Group) SetContriCurrent(contriCurrent int32) {
	g.ContriCurrent = contriCurrent
	g.IsDirty = true
}

func (g *Group) SetContriLevelUp(contriLevelUp int32) {
	g.ContriLevelUp = contriLevelUp
	g.IsDirty = true
}

type GroupMember struct {
	UserId      string
	Name        string
	Level       int32
	Power       int32
	ContriToday int32
	ContriTotal int32
	Job         int32
	OffLineTime int32

	IsDirty bool
}

func NewGroupMember() *GroupMember {
	member := new(GroupMember)
	member.IsDirty = true
	return member
}

func (m *GroupMember) SetUserId(userId string) {
	m.UserId = userId
	m.IsDirty = true
}

func (m *GroupMember) SetName(name string) {
	m.Name = name
}

func (m *GroupMember) SetLevel(level int32) {
	m.Level = level
}

func (m *GroupMember) SetContriToday(contriToday int32) {
	m.ContriToday = contriToday
	m.IsDirty = true
}

func (m *GroupMember) SetContriTotal(contriTotal int32) {
	m.ContriTotal = contriTotal
	m.IsDirty = true
}

func (m *GroupMember) SetJob(job int32) {
	m.Job = job
	m.IsDirty = true
}
