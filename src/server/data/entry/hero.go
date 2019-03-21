package entry

import "fmt"

const (
	HeroTypeStrength     int32 = 1
	HeroTypeAgility      int32 = 2
	HeroRandomDiamondErr int32 = 3
)

type Hero struct {
	Id               int32
	Name             string
	Level            int32
	Type             int32
	Strength         int32
	StrengthStep     int32
	Agility          int32
	AgilityStep      int32
	Intelligence     int32
	IntelligenceStep int32
	Armor            int32
	AttackMin        int32
	AttackMax        int32
	Blood            int32
	SkillIds         []int32
	Skills           []Skill
}

func (h *Hero) String() string {
	return fmt.Sprintf("\n{Id : %d, Name : %s, Level : %d, Type : %d, Strength : %d(+%d), Agility : %d(+%d), Intelligence : %d(+%d), Armor : %d, Attack : (%d~%d), Blood : %d, SkillIds : %v, Skills : %v }",
		h.Id, h.Name, h.Level, h.Type, h.Strength, h.StrengthStep, h.Agility, h.AgilityStep, h.Intelligence, h.IntelligenceStep, h.Armor, h.AttackMin, h.AttackMax, h.Blood, h.SkillIds, h.Skills)
}
