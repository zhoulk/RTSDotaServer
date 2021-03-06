package entry

import (
	"math"

	"github.com/name5566/leaf/log"
)

/*
魔法箭

主动技能

等级1:造成100点的伤害，晕眩1.45秒。 施法间隔:10秒 施法消耗:110点魔法
等级2:造成175点的伤害，晕眩1.55秒。 施法间隔:10秒 施法消耗:120点魔法
等级3:造成250点的伤害，晕眩1.65秒。 施法间隔:10秒 施法消耗:130点魔法
等级4:造成325点的伤害，晕眩1.75秒。 施法间隔:10秒 施法消耗:140点魔法"

向一个敌方单位射出魔法箭，眩晕并造成伤害。

*/

type SkillAct006 struct {
}

func (a SkillAct006) IsCoolDown(timer int32, skill *Skill) bool {
	if timer%10000 == 0 {
		return true
	}
	return false
}

func (a SkillAct006) Attack(timer int32, skill *Skill, from *Hero, heros []*Hero) bool {
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
		log.Fatal("[SkillAct006 Attack] Error skill.Level = %d", skill.Level)
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
