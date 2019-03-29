package entry

import (
	"github.com/name5566/leaf/log"
)

type SkillAct interface {
	IsCoolDown(timer int32, skill *Skill) bool
	Attack(timer int32, skill *Skill, from *Hero, heros []*Hero) bool
}

func NewSkillAct(skill *Skill) SkillAct {
	var act SkillAct
	switch skill.Id {
	case 1:
		act = new(SkillAct001)
		break
	case 2:
		act = new(SkillAct002)
		break
	case 3:
		act = new(SkillAct003)
		break
	case 4:
		act = new(SkillAct004)
		break
	case 6:
		act = new(SkillAct006)
		break
	case 7:
		act = new(SkillAct007)
		break
	case 8:
		act = new(SkillAct008)
		break
	case 9:
		act = new(SkillAct009)
		break
	default:
		log.Fatal("[SkillAct ] Error skill.Id = %d", skill.Id)
		break
	}
	return act
}
