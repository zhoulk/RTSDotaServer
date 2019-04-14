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
}
