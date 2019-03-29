package entry

import (
	"math"

	"github.com/name5566/leaf/log"
)

/*
恐怖波动

主动技能

魔法消耗  40
持续时间：20秒。
冷却时间  15
最大距离  1400
作用范围  350
目标单位  所有敌方单位

等级1 - 减少2点的护甲并造成30点伤害。
等级2 - 减少3点的护甲并造成50点伤害。
等级3 - 减少4点的护甲并造成70点伤害。
等级4 - 减少5点的护甲并造成90点伤害。"

复仇之魂放出邪恶的嚎叫，削弱敌人的护甲并打开经过路径的视野。

*/

type SkillAct007 struct {
}

func (a SkillAct007) IsCoolDown(timer int32, skill *Skill) bool {
	if timer%15000 == 0 {
		return true
	}
	return false
}

func (a SkillAct007) Attack(timer int32, skill *Skill, from *Hero, heros []*Hero) bool {
	selfGroup := from.Group
	var expendMP int32 = 40
	var effectArmor int32 = 0
	var effectBlood int32 = 0
	var reduceArmorDuration int32 = 20000
	switch skill.Level {
	case 1:
		effectArmor = 200
		effectBlood = 30
		break
	case 2:
		effectArmor = 300
		effectBlood = 50
		break
	case 3:
		effectArmor = 400
		effectBlood = 70
		break
	case 4:
		effectArmor = 500
		effectBlood = 90
		break
	default:
		log.Fatal("[SkillAct007 Attack] Error skill.Level = %d", skill.Level)
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
		buff.Duration = reduceArmorDuration
		buff.Type = Buff_ReduceArmor
		buff.Value = make([]int32, 0)
		buff.Value = append(buff.Value, effectArmor)
		attackHero.AddBuff(buff)

		return true
	}

	return false
}
