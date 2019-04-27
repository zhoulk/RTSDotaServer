package entry

import (
	"fmt"
)

const (
	SkillTypeActive  int32 = 1
	SkillTypePassive int32 = 2
)

type Skill struct {
	Id        int32
	Name      string
	Level     int32
	Type      int32
	LevelDesc []string
	Desc      string
	CoolTime  []int32
	IsOpen    bool
	HeroId    string
	SkillId   string

	act SkillAct

	IsDirty bool
}

func NewSkill() *Skill {
	skill := new(Skill)
	skill.IsDirty = true
	return skill
}

func (s *Skill) SetHeroId(heroId string) {
	s.HeroId = heroId
	s.IsDirty = true
}

func (s *Skill) SetSkillId(skillId string) {
	s.SkillId = skillId
	s.IsDirty = true
}

func (h *Skill) String() string {
	return fmt.Sprintf("\n{Id : %d, Name : %s, Level : %d, Type : %d, IsOpen : %v, desc : %s}",
		h.Id, h.Name, h.Level, h.Type, h.IsOpen, "" /*h.Desc*/)
}

func (h *Skill) Attack(timer int32, from *Hero, heros []*Hero) bool {
	if !h.IsCoolDown(timer) {
		return false
	}

	return h.Act().Attack(timer, h, from, heros)
}

func (h *Skill) IsCoolDown(timer int32) bool {
	return h.Act().IsCoolDown(timer, h)
}

func (h *Skill) Act() SkillAct {
	if h.act == nil {
		h.act = NewSkillAct(h)
	}
	return h.act
}
