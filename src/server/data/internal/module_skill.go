package internal

import "server/data/entry"

func (m *Module) FindASkill(skillId int32) *entry.Skill {
	skills := m.AllSkills()
	for _, skill := range skills {
		if skillId == skill.Id {
			return skill
		}
	}
	return nil
}

func (m *Module) AllSkills() []*entry.Skill {
	return m.skills
}

func (m *Module) FindHeroSkill(player *entry.Player, skillId string) (*entry.Hero, *entry.Skill) {
	heros := m.AllOwnHeros(player)
	for _, hero := range heros {
		if hero.Skills != nil {
			for _, skill := range hero.Skills {
				if skill.SkillId == skillId {
					return hero, skill
				}
			}
		}
	}

	return nil, nil
}

func (m *Module) SkillUpgrade(hero *entry.Hero, skill *entry.Skill) {
	skill.Level += 1
	skill.IsOpen = true
	hero.SkillPoint -= 1
}
