package entry

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
	return false
}
