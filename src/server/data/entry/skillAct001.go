package entry

import (
	"math"

	"github.com/name5566/leaf/log"
)

/*
死亡缠绕

主动技能

速度 5

等级 1 - 以自身75点生命的代价，伤害/治疗一个敌方/友方单位100点的生命。
等级 2 - 以自身100点生命的代价，伤害/治疗一个敌方/友方单位150点的生命。
等级 3 - 以自身125点生命的代价，伤害/治疗一个敌方/友方单位200点的生命。
等级 4 - 以自身150点生命的代价，伤害/治疗一个敌方/友方单位250点的生命。

以自身生命为祭，通过死亡缠绕伤害/治疗一个敌方/友方单位。
*/

type SkillAct001 struct {
}

func (a SkillAct001) IsCoolDown(timer int32, skill *Skill) bool {
	if timer%5000 == 0 {
		return true
	}
	return false
}

func (a SkillAct001) Attack(timer int32, skill *Skill, from *Hero, heros []*Hero) bool {
	selfGroup := from.Group
	var expendBlood int32 = 0
	var effectBlood int32 = 0
	switch skill.Level {
	case 1:
		expendBlood = 75
		effectBlood = 100
		break
	case 2:
		expendBlood = 100
		effectBlood = 150
		break
	case 3:
		expendBlood = 125
		effectBlood = 200
		break
	case 4:
		expendBlood = 150
		effectBlood = 250
		break
	default:
		log.Fatal("[SkillAct001 Attack] Error skill.Level = %d", skill.Level)
		break
	}

	var targetHero *Hero
	var attackHero *Hero
	var selfMinBlood int32 = math.MaxInt32
	var otherMinBlood int32 = math.MaxInt32
	for _, hero := range heros {
		// 优先补血
		if hero.Group == selfGroup && hero.Id != from.Id && hero.MaxBlood-hero.Blood >= effectBlood {
			if hero.Blood < selfMinBlood {
				selfMinBlood = hero.Blood
				targetHero = hero
			}
		}
		if hero.Group != selfGroup {
			if hero.Blood < otherMinBlood {
				otherMinBlood = hero.Blood
				attackHero = hero
			}
		}
	}
	// 补血
	if targetHero != nil {
		from.Blood -= expendBlood
		targetHero.Blood += effectBlood

		log.Debug("[Attack ] %d %s(%d) 使用技能 %v(%d) lv:%d 给 %s(%d) 加血 %d, 消耗 %d HP", timer, from.Name, from.Blood, skill.Name, skill.Id, skill.Level, targetHero.Name, targetHero.Blood, effectBlood, expendBlood)

		return true
	}
	// 攻击
	if attackHero != nil {
		from.Blood -= expendBlood
		attackHero.Blood -= effectBlood

		log.Debug("[Attack ] %d %s(%d) 使用技能 %v(%d) lv:%d 对 %s(%d) 造成伤害 %d, 消耗 %d HP", timer, from.Name, from.Blood, skill.Name, skill.Id, skill.Level, attackHero.Name, attackHero.Blood, effectBlood, expendBlood)
		return true
	}
	return false
}
