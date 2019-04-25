package internal

import (
	"server/data/entry"
	"server/tool"

	"github.com/name5566/leaf/log"
)

func (m *Module) AllHeros() []*entry.Hero {
	return m.heros
}

func (m *Module) HasHero(id int32) bool {
	exist := false
	for _, define := range m.heros {
		if define.Id == id {
			exist = true
			break
		}
	}

	return exist
}

func (m *Module) AllOwnHeros(player *entry.Player) []*entry.Hero {
	if player == nil || len(player.UserId) == 0 {
		log.Error("[AllOwnHeros ] player is nil or userId length is 0")
		return nil
	}
	return m.playerHeros[player.UserId]
}

func (m *Module) FindAHero(player *entry.Player, heroId string) *entry.Hero {
	heros := m.AllOwnHeros(player)
	if heros != nil {
		for _, hero := range heros {
			if hero.HeroId == heroId {
				return hero
			}
		}
	}
	return nil
}

func (m *Module) FindAHeroAt(player *entry.Player, pos int32) *entry.Hero {
	heros := m.AllOwnHeros(player)
	if heros != nil {
		for _, hero := range heros {
			if hero.Pos == pos {
				return hero
			}
		}
	}
	return nil
}

func (m *Module) FindHeroSkills(hero *entry.Hero) []*entry.Skill {
	if hero.Skills == nil {
		skills := make([]*entry.Skill, 0)
		if hero.SkillIds != nil {
			for _, skillId := range hero.SkillIds {
				skill := m.FindASkill(skillId)
				if skill != nil {
					sk := new(entry.Skill)
					tool.DeepCopy(sk, skill)
					sk.HeroId = hero.HeroId
					sk.SkillId = tool.UniqueId()
					skills = append(skills, sk)
				} else {
					log.Error("[FindHeroSkills ] skill is not exist , skillId = %v", skillId)
				}
			}
		}
		hero.Skills = skills
	}
	return hero.Skills
}

func (m *Module) RemoveHero(player *entry.Player, heroId string) *entry.Hero {
	heros := m.AllOwnHeros(player)
	var oldHero *entry.Hero
	if heros != nil {
		target := heros[:0]
		for _, hero := range heros {
			if hero.HeroId != heroId {
				target = append(target, hero)
			} else {
				oldHero = hero
			}
		}
	}
	return oldHero
}

func (m *Module) UnSelectHero(player *entry.Player, hero *entry.Hero) {
	hero.IsSelect = false
	hero.Pos = 0
}

func (m *Module) SelectHero(player *entry.Player, hero *entry.Hero, pos int32) {
	hero.IsSelect = true
	hero.Pos = pos
}

func (m *Module) SelectHeroIds(player *entry.Player) []string {
	heros := m.AllOwnHeros(player)
	heroIds := make([]string, 0)
	if heros != nil {
		for _, hero := range heros {
			if hero.IsSelect {
				heroIds = append(heroIds, hero.HeroId)
			}
		}
	}
	return heroIds
}

func (m *Module) SelectHeros(player *entry.Player) []*entry.Hero {
	heros := m.AllOwnHeros(player)
	selelctHeros := make([]*entry.Hero, 0)
	if heros != nil {
		for _, hero := range heros {
			if hero.IsSelect {
				selelctHeros = append(selelctHeros, hero)
			}
		}
	}
	return selelctHeros
}
