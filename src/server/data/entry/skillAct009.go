package entry

import (
	"math"

	"github.com/name5566/leaf/log"
)

/*
移形换位

主动技能

45

"等级 1 - 施法距离700。
等级 2 - 施法距离950。
等级 3 - 施法距离1200。"

"与一个目标英雄瞬间交换位置，无论敌我。移形换位打断目标的持续施法。 可用神杖升级。
如果你将目标移形换位到一个不可到达的区域，该目标将拥有5秒无视地形的能力
神杖效果：使大招CD时间由45秒变为10秒"

*/

type SkillAct009 struct {
}

func (a SkillAct009) IsCoolDown(timer int32, skill *Skill) bool {
	if timer%10000 == 0 {
		return true
	}
	return false
}

func (a SkillAct009) Attack(timer int32, skill *Skill, from *Hero, heros []*Hero) bool {
	selfGroup := from.Group
	var expendMP int32 = 0
	var effectBlood int32 = 0
	var dizzyDuration int32 = 0
	switch skill.Level {
	case 1:
		expendMP = 110
		effectBlood = 100
		dizzyDuration = 1450
		break
	case 2:
		expendMP = 120
		effectBlood = 175
		dizzyDuration = 1550
		break
	case 3:
		expendMP = 130
		effectBlood = 250
		dizzyDuration = 1650
		break
	case 4:
		expendMP = 140
		effectBlood = 325
		dizzyDuration = 1750
		break
	default:
		log.Fatal("[SkillAct009 Attack] Error skill.Level = %d", skill.Level)
		break
	}

	if from.MP < expendMP {
		return false
	}

	var attackHero *Hero
	var otherMinBlood int32 = math.MaxInt32
	for _, hero := range heros {
		if hero.Group != selfGroup {
			if hero.Blood < otherMinBlood {
				otherMinBlood = hero.Blood
				attackHero = hero
			}
		}
	}

	// 攻击
	if attackHero != nil {
		from.MP -= expendMP
		attackHero.Blood -= effectBlood

		log.Debug("[Attack ] %d %s(%d) 使用技能 %v(%d) lv:%d 对 %s(%d) 造成伤害 %d, 消耗 %d MP", timer, from.Name, from.Blood, skill.Name, skill.Id, skill.Level, attackHero.Name, attackHero.Blood, effectBlood, expendMP)

		buff := new(Buff)
		buff.Start = timer
		buff.Duration = dizzyDuration
		buff.Type = Buff_Dizzy
		attackHero.AddBuff(buff)

		return true
	}

	return false
}
