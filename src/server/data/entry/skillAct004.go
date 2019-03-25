package entry

import "github.com/name5566/leaf/log"

/*
回光返照

主动技能

60/50/40秒

等级 1 - 持续4秒。
等级 2 - 持续5秒。
等级 3 - 持续6秒。

启动时，移除身上负面的魔法效果，期间所有受到的伤害转而治疗你。如果你当前的生命值低于400而技能不在CD过程中，则技能会自动启动。

*/

type SkillAct004 struct {
}

func (a SkillAct004) IsCoolDown(timer int32, skill *Skill) bool {
	var offset int32 = 1
	switch skill.Level {
	case 1:
		offset = 60000
		break
	case 2:
		offset = 50000
		break
	case 3:
		offset = 40000
		break
	default:
		log.Fatal("[SkillAct004 IsCoolDown ] Error skill.Level = %d", skill.Level)
	}
	if timer%offset == 0 {
		return true
	}
	return false
}

func (a SkillAct004) Attack(timer int32, skill *Skill, from *Hero, heros []*Hero) bool {
	return false
}
