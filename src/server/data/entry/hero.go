package entry

import "fmt"

type Hero struct {
	Id           string
	Name         string
	Level        int32
	Strength     int32
	Agility      int32
	Intelligence int32
	Armor        int32
	Attack       int32
	Blood        int32
	SkillIds     []string
	Skills       []Skill
}

func (h *Hero) String() string {
	return fmt.Sprintf("{Id : %s, Name : %s, Level : %d, Strength : %d, Agility : %d, Intelligence : %d, Armor : %d, Attack : %d, Blood : %d, Skills : %v }",
		h.Id, h.Name, h.Level, h.Strength, h.Agility, h.Intelligence, h.Armor, h.Attack, h.Blood, h.Skills)
}
