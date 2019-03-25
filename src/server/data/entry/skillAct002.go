package entry

import (
	"github.com/name5566/leaf/log"
)

/*
无光之盾

主动技能

12/10/8/6秒

持续15s

等级 1 - 总共能吸收110点的伤害。
等级 2 - 总共能吸收140点的伤害。
等级 3 - 总共能吸收170点的伤害。
等级 4 - 总共能吸收200点的伤害。

用黑暗能量创造一个盾牌来保护友方的单位，在盾牌消失前吸收一定量的伤害。在盾牌被摧毁或持续时间到后，会给周围500范围内的敌方单位造成伤害。施放时移除目标身上的负面魔法效果。

*/

type SkillAct002 struct {
}

func (a SkillAct002) IsCoolDown(timer int32, skill *Skill) bool {
	var offset int32 = 1
	switch skill.Level {
	case 1:
		offset = 12000
		break
	case 2:
		offset = 10000
		break
	case 3:
		offset = 8000
		break
	case 4:
		offset = 6000
		break
	default:
		log.Fatal("[SkillAct002 IsCoolDown ] Error skill.Level = %d", skill.Level)
	}
	if timer%offset == 0 {
		return true
	}
	return false
}

func (a SkillAct002) Attack(timer int32, skill *Skill, from *Hero, heros []*Hero) bool {
	return false
}
