package entry

type Player struct {
	UserId   string
	Account  string
	Password string
	Name     string

	BaseInfo *BaseInfo
}

type BaseInfo struct {
	Gold    int32
	Diamond int32
	Exp     int32
	Power   int32
	Level   int32
}

type ExtendInfo struct {
}
