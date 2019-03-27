package internal

import (
	"server/data"
	"server/data/entry"
	"server/define"
	"server/tool"
	"sort"

	"github.com/name5566/leaf/log"
)

func fightGuanKa(player *entry.Player, gk *entry.GuanKa) bool {
	isWin := false

	// 获取己方英雄
	selfHeros := data.Module.SelectHeros(player)
	log.Debug("[fight ] selfHeros = %v", selfHeros)
	// 获取对方英雄
	otherHeros := gk.Heros
	log.Debug("[fight ] otherHeros = %v", otherHeros)

	var timer int32 = 0
	fightHeros := make([]*entry.Hero, 0)

	for _, hero := range selfHeros {
		h := new(entry.Hero)
		tool.DeepCopy(h, hero)
		h.Group = 1
		h.BAT = define.FIGHT_BASE_ATTACK_TIME
		fightHeros = append(fightHeros, h)
	}

	for _, hero := range otherHeros {
		h := new(entry.Hero)
		tool.DeepCopy(h, hero)
		h.Group = 2
		h.BAT = define.FIGHT_BASE_ATTACK_TIME
		fightHeros = append(fightHeros, h)
	}

	// 装备属性生效
	EquipEffect(fightHeros)
	// 被动技能生效
	PassiveEffect(fightHeros)
	for {
		// 10Hz逻辑帧加速运算   100ms
		timer += define.FIGHT_LOGIC_FRAMGE_RATE
		// 按照BAT排序
		sort.Sort(entry.SortByBAT(fightHeros))
		// 主动技能生效 or 普通攻击
		for _, hero := range fightHeros {
			isAttack := false
			for _, skill := range hero.Skills {
				if skill.Type == entry.SkillTypeActive && skill.IsOpen {
					if skill.Attack(timer, hero, fightHeros) {
						isAttack = true
					}
				}
			}
			if !isAttack {
				// 普通攻击
				hero.NormalAttack(timer, fightHeros)
			}
		}

		// 判断死亡
		target := fightHeros[:0]
		anyOneDie := false
		for _, hero := range fightHeros {
			if hero.Blood > 0 {
				target = append(target, hero)
			} else {
				anyOneDie = true
			}
		}
		fightHeros = target

		if anyOneDie {
			// 装备属性生效
			EquipEffect(fightHeros)
			// 被动技能生效
			PassiveEffect(fightHeros)
		}
		// 判断结束
		// 只存活一组 战斗结束
		res := make(map[int32]bool)
		for _, hero := range fightHeros {
			res[hero.Group] = true
		}
		if len(res) <= 1 {
			if res[selfHeros[0].Group] {
				isWin = true
			}
			break
		}
		if timer >= define.FIGHT_MAX_DURATION {
			break
		}
	}

	return isWin
}

func EquipEffect(fightHeros []*entry.Hero) {
	log.Debug("[fight ] equip effect begin ")
	for _, hero := range fightHeros {
		if hero.Equips != nil {
			for _, equip := range hero.Equips {
				log.Debug("[fight ] equip effect %v ", equip)
			}
		}
	}
	log.Debug("[fight ] equip effect end ")
}

func PassiveEffect(fightHeros []*entry.Hero) {
	log.Debug("[fight ] skill effect begin ")
	for _, hero := range fightHeros {
		if hero.Skills != nil {
			for _, skill := range hero.Skills {
				if skill.Type == entry.SkillTypePassive && skill.IsOpen {
					log.Debug("[fight ] skill effect %v ", skill)
				}
			}
		}
	}
	log.Debug("[fight ] skill effect end ")
}
