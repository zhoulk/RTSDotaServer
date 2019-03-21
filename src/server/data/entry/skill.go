package entry

import "fmt"

const (
	SkillTypeActive  int32 = 1
	SkillTypePassive int32 = 2
)

type Skill struct {
	Id    int32
	Name  string
	Level int32
	Type  int32
	Desc  string
}

func (h *Skill) String() string {
	return fmt.Sprintf("\n{Id : %d, Name : %s, Level : %d, Type : %d, desc : %s}",
		h.Id, h.Name, h.Level, h.Type, "" /*h.Desc*/)
}
