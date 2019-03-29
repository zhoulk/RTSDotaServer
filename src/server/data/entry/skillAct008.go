package entry

import (
	"math"

	"github.com/name5566/leaf/log"
)

/*
复仇光环

被动技能

"等级 1 - 增加12%的基础攻击力。
等级 2 - 增加20%的基础攻击力。
等级 3 - 增加28%的基础攻击力。
等级 4 - 增加36%的基础攻击力。"

复仇之魂的存在提高了附近友方单位的物理攻击力。

*/

type SkillAct008 struct {
}

func (a SkillAct008) IsCoolDown(timer int32, skill *Skill) bool {
	if timer%10000 == 0 {
		return true
	}
	return false
}

func (a SkillAct008) Attack(timer int32, skill *Skill, from *Hero, heros []*Hero) bool {
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
		log.Fatal("[SkillAct008 Attack] Error skill.Level = %d", skill.Level)
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
