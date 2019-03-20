package entry

import "fmt"

type Skill struct {
	Id    string
	Name  string
	Level int32
	Desc  string
}

func (h *Skill) String() string {
	return fmt.Sprintf("{Id : %s, Name : %s, Level : %d, desc : %s}",
		h.Id, h.Name, h.Level, h.Desc)
}
