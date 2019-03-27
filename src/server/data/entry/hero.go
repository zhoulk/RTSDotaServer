package entry

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/name5566/leaf/log"
)

const (
	HeroTypeStrength     int32 = 1
	HeroTypeAgility      int32 = 2
	HeroRandomDiamondErr int32 = 3
)

type Hero struct {
	Id               int32
	Name             string
	Level            int32
	Type             int32
	Strength         int32
	StrengthStep     int32
	Agility          int32
	AgilityStep      int32
	Intelligence     int32
	IntelligenceStep int32
	Armor            int32
	AttackMin        int32
	AttackMax        int32
	Blood            int32
	MaxBlood         int32
	MP               int32
	MaxMP            int32
	SkillIds         []int32
	Skills           []*Skill
	EquipId          []int32
	Equips           []*Item
	HeroId           string
	PlayerId         string
	IsSelect         bool
	Pos              int32
	SkillPoint       int32
	Exp              int32
	LevelUpExp       int32

	BAT   int32
	Group int32
}

func (h *Hero) String() string {
	return fmt.Sprintf("\n{Id : %d, Name : %s, Level : %d, IsSelect :%v, Pos : %d, Type : %d, Strength : %d(+%d), Agility : %d(+%d), Intelligence : %d(+%d), Armor : %d, Attack : (%d~%d), Blood : (%d / %d), MP : (%d / %d), BAT : %d, SkillIds : %v, Skills : %v }",
		h.Id, h.Name, h.Level, h.IsSelect, h.Pos, h.Type, h.Strength, h.StrengthStep, h.Agility, h.AgilityStep, h.Intelligence, h.IntelligenceStep, h.Armor, h.AttackMin, h.AttackMax, h.Blood, h.MaxBlood, h.MP, h.MaxMP, h.BAT, h.SkillIds, h.Skills)
}

func (h *Hero) NormalAttack(timer int32, heros []*Hero) {
	if !h.CanAttack(timer) {
		return
	}

	var attackHero *Hero
	var otherMinBlood int32 = math.MaxInt32
	for _, hero := range heros {
		if hero.Group != h.Group {
			if hero.Blood < otherMinBlood {
				otherMinBlood = hero.Blood
				attackHero = hero
			}
		}
	}

	var effectBlood int32 = 0
	if attackHero != nil {
		selfAttack := h.AttackMin + rand.Int31n(h.AttackMax-h.AttackMin)
		if attackHero.Armor >= 0 {
			reduce := float32(attackHero.Armor/100*6) / float32(100+attackHero.Armor/100*6)
			effectBlood = int32(float32(selfAttack) * (float32(1) - reduce) / 100)
		} else {
			deeper := float64(2) - math.Pow(0.94, float64(attackHero.Armor/100))
			effectBlood = int32(float64(selfAttack) * (float64(1) + deeper) / 100)
		}

		attackHero.Blood -= effectBlood
		log.Debug("[Attack ] %d %s(%d) 攻击 %s(%d) 造成 %d 点伤害", timer, h.Name, h.Blood, attackHero.Name, attackHero.Blood, effectBlood)
	}
}

func (h *Hero) CanAttack(timer int32) bool {
	if h.BAT > 0 {
		if timer%h.BAT == 0 {
			return true
		}
	}
	return false
}

func (h *Hero) EffectByEarn(earn *Earn) {
	h.Exp += earn.HeroExp
	for {
		if h.Exp < h.LevelUpExp {
			break
		}
		if h.Level+1 >= int32(len(heroExpList)) {
			break
		}
		h.levelUp()
	}
}

func (h *Hero) levelUp() {
	h.Exp -= h.LevelUpExp
	h.Level += 1
	h.SkillPoint += 1
	h.LevelUpExp = heroExpList[h.Level]
}

type SortByBAT []*Hero

func (a SortByBAT) Len() int {
	return len(a)
}
func (a SortByBAT) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a SortByBAT) Less(i, j int) bool {
	if a[i].BAT == a[j].BAT {
		if a[i].Agility == a[j].Agility {
			if a[i].Level == a[j].Level {
				return a[i].Id < a[j].Id
			}
			return a[i].Level > a[j].Level
		}
		return a[i].Agility > a[j].Agility
	}
	return a[i].BAT < a[j].BAT
}
