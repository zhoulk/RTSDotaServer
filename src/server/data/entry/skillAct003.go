package entry

/*
霜之哀伤

被动技能

等级 1 - 降低目标5%/5%攻击/移动速度，提升攻击者10%/15%攻击/移动速度。
等级 2 - 降低目标10%/10%攻击/移动速度，提升攻击者20%/15%攻击/移动速度。
等级 3 - 降低目标15%/15%攻击/移动速度，提升攻击者30%/15%攻击/移动速度。
等级 4 - 降低目标20%/20%攻击/移动速度，提升攻击者40%/15%攻击/移动速度。

地狱领主使用传说之剑霜之哀伤的极寒力量使目标减速，持续2.5秒。任何攻击被减速目标的单位得到速度的提升，持续4.5秒


*/

type SkillAct003 struct {
}

func (a SkillAct003) IsCoolDown(timer int32, skill *Skill) bool {
	return true
}

func (a SkillAct003) Attack(timer int32, skill *Skill, from *Hero, heros []*Hero) bool {
	return false
}
